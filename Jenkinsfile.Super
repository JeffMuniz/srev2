http://machina:8081/fature/fature-auditado/sync-capeante/blob/master/Jenkinsfile#L497


#!/usr/bin/env groovy

versao = "1.0.${env.BUILD_ID}"
publish = true
def isStartedByUser = true
def FAILURE_REASON = ""

pipeline {
    agent { node { label 'dotnet-31'} }
    environment {
        //FAILURE_REASON=""
        PRJ_PATH='.'
        PUBLISH_PATH='app'
        PRJ_NAME='machina.SincronizadorCapeante.sln'
        APPLICATION_NAME='fature-auditado-sync-capeante'
        OPENSHIFT_PROJECT_NAME='fature-auditado'
        TEST_PATH='./tests/machina.Sincronizador.Capeante.Tests'

        //CONFIGMAP_DEV="src/machina.Queue.MotorContingencia.Cassi.Worker/appSettings-dev.json"
        //CONFIGMAP_HML="src/machina.Queue.MotorContingencia.Cassi.Worker/appSettings-hml.json"
        //CONFIGMAP_PRD="src/machina.Queue.MotorContingencia.Cassi.Worker/appSettings-prd.json"
        //CONFIGMAP_NAME="appsettings-configmap"
        //CONFIGMAP_VOLUME="v1-appsettings"

        SECRET_VOLUME="v1-appsettings"
        SECRET_NAME="secret-appsettings-sync-capeante"
        SECRET_KEY="appsettings.json"
        SECRET_FILENAME="src/machina.Sincronizador.Capeante/appsettings.json"

        SECRET_SENHA_DATIVA="senha-dativa"

        SONAR_PATH='/home/jenkins/.dotnet/tools'       
        SONAR_SERVER='http://sonarqube.machina.local'
        NEXUS_PKG_GROUP='fature-auditado'

        TESTS_RESULTS="${WORKSPACE}/.tests-results/results.xml"
        OPENCOVER_RESULTS="${WORKSPACE}/.opencover-results/results.xml"

        FORTIFY_PROJECT_NAME='sync-capeante'

        SONAR_PROJECT_URL='http://sonarqube.machina.local/dashboard?id=sync-capeante'
        CHAT_WEBHOOK='https://chat.googleapis.com/v1/spaces/AAAAAevr2i4/messages?key=AIzaS...'
    }
    options {
      gitLabConnection('Gitlab')
      disableConcurrentBuilds()
    }
    triggers {
        gitlab(
            triggerOnPush: true,
            triggerOnMergeRequest: true,
            branchFilterType: 'All',
            secretToken: 'f7e206')
    }
    stages 
    {
        stage ( DEV ) {
            when {
                allOf {
                    expression{ env.deploy_environment == "dev" }
                }
            }
            stages {
                stage('Validate-Branch') {
                    steps 
                    {
                        script
                        {
                            if ( currentBuild.rawBuild.getCauses()[0].toString().contains('UserIdCause') )
                            {
                                isStartedByUser = false
                            }

                            if(isStartedByUser)
                            {
                                echo '*************************************************************************'
                                echo "QUEM FEZ COMMIT? ${env.gitlabUserName}"
                                echo '*************************************************************************'  

                                echo gitlabSourceBranch            

                                //if (!gitlabSourceBranch.startsWith('release')) 
                                //{
                                //    //currentBuild.result = 'ABORTED'
                                //    publish = false

                                    //setError(message: "Branch invalida para build na esteira!")
                                    //return
                                //}
                            } 
                            else
                            {
                                echo '*************************************************************************'
                                echo "QUEM FEZ COMMIT? BUILD MANUAL"
                                echo '*************************************************************************'       
                            }
                        }                 
                    }           
                }
                stage('pre build') {
                    steps {        
                        script {
                            branchCreated = (env.gitlabBefore == "0000000000000000000000000000000000000000")                    
                        }
                        sh 'rm -rf ${PUBLISH_PATH}'
                        sh 'mkdir ${PUBLISH_PATH}'
                    }
                }
                stage('clean') {
                    // when {
                    //     allOf {
                    //         not { equals actual: branchCreated, expected: true }
                    //         expression{ publish == true }
                    //     }
                    // }

                    steps 
                    {
                        sh 'dotnet clean ${PRJ_PATH}/${PRJ_NAME}'
                        sh 'dotnet restore -s http://nexus.machina.local/repository/nuget-group/ -s https://api.nuget.org/v3/index.json ${PRJ_PATH}/${PRJ_NAME}'
                    }
                }
                stage('sonar-config') {
                   steps 
                   {  
                        sh 'dotnet tool restore'               
                   }
                }   
                stage('Sonar-Start') {
                    steps
                    {
                        script 
                        {
                            withSonarQubeEnv('sonarqube machina')  
                            {           
                                echo "QUEM FEZ COMMIT? ${env.gitlabUserName}"
                                echo "${APPLICATION_NAME}"                                
                                sh "dotnet dotnet-sonarscanner begin /key:${APPLICATION_NAME} /name:${APPLICATION_NAME} /version:1.${env.BUILD_ID} /d:sonar.host.url=${SONAR_SERVER} /d:sonar.verbose=true /d:sonar.cs.opencover.reportsPaths=${OPENCOVER_RESULTS} /d:sonar.analysis.projeto=\"Fature-Auditado\" /d:sonar.analysis.user=\"${env.gitlabUserName}\" /d:sonar.cs.xunit.reportsPaths=\"${TESTS_RESULTS}\""                 
                            }
                        }
                    }
                }   
                stage('build') {
                    // when {
                    //     allOf {
                    //         not { equals actual: branchCreated, expected: true }
                    //         expression{ publish == true }
                    //     }
                    // }

                    steps 
                    {
                        sh 'dotnet build --no-restore -c Release ${PRJ_PATH}/${PRJ_NAME}'            
                    }
                }
                stage('Unit-Tests') {
                    steps
                    {       
                        sh "mkdir -p ${WORKSPACE}/.opencover-results"
                        sh "mkdir -p ${WORKSPACE}/.tests-results"  
                        sh 'dotnet test ${PRJ_PATH}/${TEST_PATH} --logger:"xunit;LogFilePath=${TESTS_RESULTS}" /p:CollectCoverage=true /p:CoverletOutputFormat=opencover /p:CoverletOutput=${OPENCOVER_RESULTS}'
                    }
                }
                stage('Sonar-End') {
                    steps 
                    {
                        withSonarQubeEnv('sonarqube machina')
                        {
                            sh "dotnet dotnet-sonarscanner end "
                        }
                    }
                }
                stage('Quality & Security-gates') {
                    parallel
                    {        
                        stage('check-quality-gate') 
                        {
                            steps 
                            {
                                echo 'Checking quality gate...'
                                script 
                                {
                                    timeout(time: 1, unit: 'HOURS') 
                                    {
                                        sleep(80)
                                        def qg = waitForQualityGate()
                                        if (qg.status == 'ERROR') {
                                            setError(message: "Codigo fonte nao esta de acordo com as metas de qualidade definidas no Sonar: ${qg.status}")
                                            return
                                        }
                                    }
                                }
                            }
                        }
                        /*stage('check-security-gate')
                        {
                            steps
                            {
                                script
                                {
                                    def handle = triggerRemoteJob(
                                        job: 'http://jenkins-ci.machina.local/job/RunFortify',
                                        parameters: "GitCommitID=${env.GIT_COMMIT}\nGitUrl=${env.GIT_URL}\nProjectName=${env.FORTIFY_PROJECT_NAME}\nSolutionName=${env.PRJ_NAME}\nBU=${env.JKS_PRODUCT}",
                                        blockBuildUntilComplete: true,
                                        auth: CredentialsAuth(credentials: 'svc_hpdevops')
                                    )
                                    echo 'Remote Status: ' + handle.getBuildStatus().toString()
                                }
                            }
                        }*/
                    }
                }
                stage('Publish') {
                    when {
                        allOf {
                            not { equals actual: branchCreated, expected: true }
                            expression{ publish == true }
                        }
                    }

                    steps 
                    {
                       sh 'dotnet publish --no-restore --no-build -c Release -o ./${PUBLISH_PATH} ${PRJ_PATH}/${PRJ_NAME}'
                       sh 'ls ${PUBLISH_PATH}'
                    }
                } 
                stage('nexus-upload') {
                    when {
                        allOf {
                            expression{ publish == true }
                            not { equals actual: branchCreated, expected: true }
                            expression{ env.gitlabSourceBranch ==~ /(master)/ }
                        }
                    }

                    steps {

                        sh "zip -r ${APPLICATION_NAME}-${env.BUILD_ID}.zip ${PRJ_PATH}/${PUBLISH_PATH}"

                        nexusArtifactUploader (
                            nexusVersion: 'nexus3',
                            protocol: 'http',
                            nexusUrl: 'nexus.machina.local',
                            repository: 'fature',
                            credentialsId: 'adminnexus',
                            version: "${versao}",
                            groupId: "${NEXUS_PKG_GROUP}",
                            artifacts: [
                                [artifactId: "${APPLICATION_NAME}",
                                type: 'zip',
                                classifier: '',                        
                                file: "./${APPLICATION_NAME}-${env.BUILD_ID}.zip"]
                            ]
                        )
                    }
                } 
                stage('[CD] - Create image') {
                    when {
                        allOf {
                            not { equals actual: branchCreated, expected: true }
                            expression {
                                openshift.withCluster() {
                                    openshift.withProject("${OPENSHIFT_PROJECT_NAME}-dev") {
                                        echo "USING PROJECT: ${openshift.project()}" 
                                        if(!publish)
                                            return false
                                        else  
                                            return !openshift.selector("bc", "${APPLICATION_NAME}").exists();
                                            currentBuild.result = 'SUCCESS'
                                    }
                                }
                            }
                        }
                    }
                    steps {
                        script {
                            openshift.withCluster() {                        
                                openshift.withProject("${OPENSHIFT_PROJECT_NAME}-dev") {
                                    echo "Creating new image"
                                    openshift.newBuild("--name=${APPLICATION_NAME}", "--image-stream=cicd-tools/dotnet-31-sdk-runtime-rhel7", "--binary=true", "--env-file=.s2i/environment")
                                }
                            }
                        }
                    }
                }
                stage('[CD] - Build image') {
                    when {
                        allOf {
                            not { equals actual: branchCreated, expected: true }
                            expression{ publish == true }
                        }
                    }
                    steps {
                        script {
                            openshift.withCluster() {
                                openshift.withProject("${OPENSHIFT_PROJECT_NAME}-dev") {
                                    echo "DIRETORIO PUBLICACAO -> ${PUBLISH_PATH}"
                                    def bc = openshift.selector("bc", "${APPLICATION_NAME}").startBuild("--from-dir=./${PUBLISH_PATH}", "--wait=true")
                                }
                            }
                        }
                    }
                }
                stage('[CD] - Create TAG DEV') {
                    when {
                        allOf {
                            not { equals actual: branchCreated, expected: true }
                            expression{ publish == true }
                        }
                    }
                    steps {
                        script {
                            openshift.withCluster() {
                                openshift.withProject("${OPENSHIFT_PROJECT_NAME}-dev") 
                                {
                                    echo "TAG: ${APPLICATION_NAME}:${versao} --> ${APPLICATION_NAME}:latest"
                                    openshift.tag("${APPLICATION_NAME}:latest", "${APPLICATION_NAME}:${versao}")
                                }
                            }
                        }
                    }
                }
                stage('[CD] - Deploy DEV') {
                    when {
                        allOf {
                            not { equals actual: branchCreated, expected: true }
                            expression{ publish == true }
                        }
                    }
                    steps {
                        script {
                            openshift.withCluster() {
                                openshift.withProject("${OPENSHIFT_PROJECT_NAME}-dev") 
                                {
                                    try {
                                        //echo "Configmap: ${CONFIGMAP_NAME} - ${CONFIGMAP_DEV}"
                                        //if (openshift.selector("configmap", "${CONFIGMAP_NAME}").exists()) {
                                        //    openshift.delete("configmap", "${CONFIGMAP_NAME}")
                                        //}
                                        //def cm = openshift.create("configmap", "${CONFIGMAP_NAME}", "--from-file=appSettings.json=${CONFIGMAP_DEV}")
                                        //def CONFIGMAP_KEY = cm.object().data.keySet()[0]

                                        echo "Secret: ${SECRET_NAME} - ${SECRET_FILENAME}"
                                        if (openshift.selector("secret", "${SECRET_NAME}").exists()) {
                                            if( !openshift.delete("secret", "${SECRET_NAME}") ) {
                                                setError(message: "Falha ao apagar Secret ${SECRET_NAME}")
                                                return
                                            }
                                        }
                                        def cm = openshift.create("secret", "generic", "${SECRET_NAME}", "--from-file=${SECRET_KEY}=${SECRET_FILENAME}", "--type=Opaque")

                                        def dcSelector = openshift.selector("dc", "${APPLICATION_NAME}")
                                        if ( dcSelector.exists() ) {
                                            def rm = dcSelector.rollout()
                                            def result = rm.history() // Gather history for the selected rollouts
                                            echo "DeploymentConfig history:\n${result.out}"
                                            echo "Realizando o Rollout da última versão...."
                                            rm.latest()
                                        } 
                                        else {
                                            echo "Criando Application ${APPLICATION_NAME}"
                                            def app = openshift.newApp("${APPLICATION_NAME}:latest",  "--name=${APPLICATION_NAME}")
                                            app.narrow("svc").expose();
                                            openshift.set("volume", "dc/${APPLICATION_NAME}", "--add", "--name=${SECRET_NAME}", "-t secret", "--secret-name=${SECRET_NAME}", "-m /opt/app-root/app/${SECRET_KEY}", "--sub-path=${SECRET_KEY}", "--default-mode=775")
                                            openshift.set("env", "dc/${APPLICATION_NAME}", "TZ=America/Sao_Paulo")
                                            //openshift.set("env", "dc/${APPLICATION_NAME}", "--from=secret/${SECRET_SENHA_DATIVA}")
                                            //openshift.set("volume", "dc/${APPLICATION_NAME}", "--add", "--name=${CONFIGMAP_VOLUME}", "-t configmap", "--configmap-name=${CONFIGMAP_NAME}", "-m /opt/app-root/app/${CONFIGMAP_KEY}", "--sub-path=${CONFIGMAP_KEY}", "--default-mode=775")
                                        }

                                        10.times { 
                                            echo "Aguardando POD subir... ${it}"
                                            def dc = openshift.selector("dc", "${APPLICATION_NAME}")
                                            if( dc.object().spec.replicas != dc.object().status.availableReplicas ) {
                                                sleep(10)
                                            } else {
                                                return;
                                            }
                                        }
                                        openshift.set("triggers", "dc/${APPLICATION_NAME}", "--manual")
                                    } catch (Exception e) {
                                        setError(message: "\"" + e.toString() + "\"")
                                        return
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
        stage ( HML ) {
            when {
                allOf {
                    expression{ env.deploy_environment == "hml" }
                }
            }
            stages {
                stage('[CD] - Promote HML') {
                    when {
                        allOf {
                            expression{ publish == true }
                        }
                    }
                    steps {
                        script {
                            timeout(time:10, unit:'MINUTES') {
                               input message: "Promote TAG ${env.image_tag} to HML?", ok: "Promote"
                            }

                            echo 'openshift.tag HML'
                            openshift.withCluster() {
                                openshift.withProject("${OPENSHIFT_PROJECT_NAME}-hml") {
                                    // check se image tag solicitada existe no Registry
                                    def existe_tag_registry = sh(script: "oc get is fature-auditado-sync-capeante -n fature-auditado-dev --template=\"{{range .spec.tags}}{{.name}}{{' '}}{{end}}\" | grep ${env.image_tag} | wc -l", returnStdout: true)
                                    echo "Existe TAG? -> ${existe_tag_registry}"
                                    if ( existe_tag_registry != 1 ) {
                                        autoCancelled = true
                                        setError(message: "Tag [${env.image_tag}] nao existe no Openshift Registry.")
                                        return
                                    }
                                    openshift.tag("${OPENSHIFT_PROJECT_NAME}-dev/${APPLICATION_NAME}:${env.image_tag}", "${OPENSHIFT_PROJECT_NAME}-hml/${APPLICATION_NAME}:${env.image_tag}")
                                    openshift.tag("${OPENSHIFT_PROJECT_NAME}-hml/${APPLICATION_NAME}:${env.image_tag}", "${OPENSHIFT_PROJECT_NAME}-hml/${APPLICATION_NAME}:latest")
                                }
                            }
                        }
                    }
                }
                stage('[CD] - Deploy HML') {
                    when {
                        allOf {
                            expression{ publish == true }
                        }
                    }
                    steps {
                        script {
                            openshift.withCluster() {
                                openshift.withProject("${OPENSHIFT_PROJECT_NAME}-hml") 
                                {
                                    def dcSelector = openshift.selector("dc", "${APPLICATION_NAME}")
                                    if ( dcSelector.exists() ) {
                                        echo "Apagando Deployment Config..."
                                        if ( !openshift.delete('dc', "${APPLICATION_NAME}")    ||
                                             !openshift.delete('route', "${APPLICATION_NAME}") ||
                                             !openshift.delete('service', "${APPLICATION_NAME}") ) {
                                            autoCancelled = true
                                            setError(message: "Falha ao apagar DeploymentConfig / Route / Service ${APPLICATION_NAME}")
                                            return
                                        }
                                    }

                                    echo "Criando Application ${APPLICATION_NAME}"
                                    def app = openshift.newApp("${APPLICATION_NAME}:latest",  "--name=${APPLICATION_NAME}")
                                    app.narrow("svc").expose();
                                    openshift.set("volume", "dc/${APPLICATION_NAME}", "--add", "--name=${SECRET_NAME}", "-t secret", "--secret-name=${SECRET_NAME}", "-m /opt/app-root/app/${SECRET_KEY}", "--sub-path=${SECRET_KEY}", "--default-mode=775")
                                    openshift.set("env", "dc/${APPLICATION_NAME}", "TZ=America/Sao_Paulo")

                                    10.times { 
                                        echo "Aguardando POD subir... ${it}"
                                        def dc = openshift.selector("dc", "${APPLICATION_NAME}")
                                        if( dc.object().spec.replicas != dc.object().status.availableReplicas ) {
                                            sleep(10)
                                        } else {
                                            return;
                                        }
                                    }
                                    openshift.set("triggers", "dc/${APPLICATION_NAME}", "--manual")
                                }
                            }
                        }
                    }
                }
            }
        }
        stage ( PRD ) {
            when {
                allOf {
                    expression{ env.deploy_environment == "prd" }
                }
            }
            stages {
                stage('[CD] - Promote PRD') {
                    when {
                        allOf {
                            expression{ publish == true }
                        }
                    }
                    steps {
                        script {
                            timeout(time:10, unit:'MINUTES') {
                               input message: "Promote TAG ${env.image_tag} to PRD?", ok: "Promote"

                                sh """
                                    mkdir -p ~/.docker
                                    cp /tmp/docker-config/config.json ~/.docker
                                    oc image mirror --insecure docker-registry-default.cpdh.machina.local/${OPENSHIFT_PROJECT_NAME}-hml/${APPLICATION_NAME}:${env.image_tag} docker-registry-default.cppr.machina.local/${OPENSHIFT_PROJECT_NAME}/${APPLICATION_NAME}:${env.image_tag}
                                """

                                openshift.withCluster('openshift_cluster_prod') {
                                    openshift.withProject("${OPENSHIFT_PROJECT_NAME}") {
                                        openshift.tag("${OPENSHIFT_PROJECT_NAME}/${APPLICATION_NAME}:${env.image_tag}", "${OPENSHIFT_PROJECT_NAME}/${APPLICATION_NAME}:latest")
                                    }
                                }
                            }
                        }
                    }
                }
                stage('[CD] - Deploy PRD') {
                    when {
                        allOf {
                            expression{ publish == true }
                        }
                    }
                    steps {
                        script {
                            openshift.withCluster('openshift_cluster_prod') {
                                openshift.verbose()
                                openshift.logLevel(3)
                                openshift.withProject("${OPENSHIFT_PROJECT_NAME}") 
                                {
                                    def dcSelector = openshift.selector("dc", "${APPLICATION_NAME}")
                                    if ( dcSelector.exists() ) {
                                        echo "Apagando Deployment Config..."
                                        if ( !openshift.delete('dc', "${APPLICATION_NAME}")    ||
                                             !openshift.delete('route', "${APPLICATION_NAME}") ||
                                             !openshift.delete('service', "${APPLICATION_NAME}") ) {
                                            autoCancelled = true
                                            setError(message: "Falha ao apagar DeploymentConfig / Route / Service ${APPLICATION_NAME}")
                                            return
                                        }
                                    }

                                    echo "Criando Application ${APPLICATION_NAME}"
                                    def app = openshift.newApp("${APPLICATION_NAME}:latest",  "--name=${APPLICATION_NAME}")
                                    app.narrow("svc").expose();
                                    openshift.set("volume", "dc/${APPLICATION_NAME}", "--add", "--name=${SECRET_NAME}", "-t secret", "--secret-name=${SECRET_NAME}", "-m /opt/app-root/app/${SECRET_KEY}", "--sub-path=${SECRET_KEY}", "--default-mode=775")
                                    openshift.set("env", "dc/${APPLICATION_NAME}", "TZ=America/Sao_Paulo")

                                    10.times { 
                                        echo "Aguardando POD subir... ${it}"
                                        def dc = openshift.selector("dc", "${APPLICATION_NAME}")
                                        if( dc.object().spec.replicas != dc.object().status.availableReplicas ) {
                                            sleep(10)
                                        } else {
                                            return;
                                        }
                                    }
                                    openshift.set("triggers", "dc/${APPLICATION_NAME}", "--manual")
                                }
                                openshift.verbose(false)
                            }
                        }
                    }
                }
            }
        }
    }
    post {
        always  {
            script {
                // Para lista completa de cows, acesse: http://cowsay-fature-auditado-dev.cpdh.machina.local/fun.pl?list
                // Para uso no Google Chat, inclua a opção uso=chat
                //def cows="ghostbusters,cower,eyes,hellokitty,dragon,daemon,bud-frogs,stimpy,default,apt,elephant-in-snake,koala,www,stegosaurus,luke-koala,calvin,tux";
                def cows="bud-frogs,default,apt,elephant-in-snake,luke-koala";
                def msg = "${JOB_BASE_NAME}%20-%20${currentBuild.currentResult}"

                def response = httpRequest "http://cowsay-fature-auditado-dev.cpdh.machina.local/fun.pl?addmsg=${msg}&uso=chat&cows=${cows}"

                def QUOTE = response.content ?: ":-(";
                def deployed_env = "undefined"

                switch(env.deploy_environment) {
                  case "dev":
                    deployed_env = "DESENVOLVIMENTO"
                  break;
                  case "hml":
                    deployed_env = "HOMOLOGACAO"
                  break;
                  case "prd":
                    deployed_env = "PRODUCAO"
                  break;
                  default:
                    deployed_env = "undefined"
                  break;
                }

                def msg_tag = ""
                if (!env.image_tag || env.image_tag == "Desabilitado") {
                    msg_tag = "${versao}"
                }

                def link_console =""
                if ( env.FAILURE_REASON || currentBuild.result == "FAILURE" ) {
                    link_console = "${BUILD_URL}console"
                }
                googlechatnotification (
                    url:"${CHAT_WEBHOOK}",
                    message:    
                                "*Environment: ${deployed_env}* \n" +
                                "*TAG:* ${msg_tag} \n" +
                                "*Author:* ${env.gitlabUserName} \n" +                                           
                                "*Build:* ${env.BUILD_ID} \n" +
                                "*Duration:* ${currentBuild.durationString} \n" +
                                "<${SONAR_PROJECT_URL}|*SONAR REPORT*> \n" +
                                ( link_console ? "<${link_console}|*Jenkins Failure Report*> \n" : "" ) +
                                "```\n${QUOTE}``` \n",
                    notifyFailure: 'true',
                    notifySuccess: 'true',
                    notifyAborted: 'true',
                    sameThreadNotification: 'true'
                    )
            }
        }

        failure {
            updateGitlabCommitStatus name: 'build', state: 'failed'

        }

        success {
            updateGitlabCommitStatus name: 'build', state: 'success'
        }
    }
}


def setError(Map args) {
    env.FAILURE_REASON = args.message
    currentBuild.result = args.buildResult ?: 'SUCCESS'
    error(args.message)
}
