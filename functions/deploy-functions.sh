#/bin/sh

############################################
## variables
GCLOUD_PROJECT="dev-dx"

############################################
## deploy Cloud Functions
cd ./command-parser
gcloud functions deploy ToilEliminatorCommandParser --runtime go111 --trigger-http --region asia-northeast1 --project ${GCLOUD_PROJECT}

# cd ../executor
# gcloud functions deploy ToilEliminatorCommandParser --runtime go111 --trigger-http --region asia-northeast1 --project ${GCLOUD_PROJECT}