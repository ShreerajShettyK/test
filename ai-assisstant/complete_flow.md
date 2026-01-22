AI-Assistant Flow:

1) for user when he asks a question to ai-assistant:

UI frontend----> Ui backend (authorization and profile token will be checked)------> chat-recorder to fetch conversations (past questions) -----> when user enters the query ----> question will be taken from ui to ui 's backend ----> UI backend ------> Secure-shield (USES gemini model) (blocks harmfull question) (true or false) ------->
if attack detected is false :
-----> go to learning-assistant service
                                        ------> (Skip if new question in new thread)
                                             chat recorder --- to fetch chats and conversations (for context when calling cache and proxy) 


                                        ------> cache service   (return the already asked question's answer based on similarity check)


                                        -------> Proxy service - talks to llms azure openai, aws bedrock, circuit, apim and gemini. Also streams back the response in chunks
                                                                                
                                        ---->
                                            citation (gets sources related to the questions asked like videos, podcasts, etc from cisco db) ---------------> embeddings of all cisco data

                                        ----->
                                             related questions---- fecthes all the questions related to the question the user asked

                                        ------>
                                             topic mapper---- creates graph which gives a visual representation of all the topics related to the one which was asked

                                        -------> Embedding service (used by cache, citation, toppic mapper) uses half vctor (not full vector)

if attack detected is true :
---------------------------------> exit and retrun  template response telling that "the question contains harmful message"



2) for users who asks questions related to training questions

 Training recommender (for training related questions)

3) Testing

 gen-ai-hammer -------- used for testing perfomrance of ai powered applications using vegeta and hey packages                                       

4) Used by analytics team so that they get to know 

 gen-ai-reports-service --- used by analytics
 takes all the input qa pairs and in batches send it to azure model and groups it to categories defined in the prompt and excel will be uploaded to s3




EXAM's question GENERATOR (new feature) FLOW:


1) for the admin (subject matter expert)---------------------------------kafka
ui backend calls for admin ops below//

                task.preflight passed                           task.questionsgenerated
 exam assistant ----------------------> question generator -----------------------------> secure-shield 
 
 task.securitypassed                                           task.readyforReview
 -----------------------------> sme_router(stores to db) -------------------------------> exam assistant


 then after all the questions are ready and generated. The exam assistant will call the sme_router for displaying the questions. Till then it will be in progress state.


Tech stack:

Nextjs , nodejs

Amazon elastic cache for redis
Amazon MSK (messaging system for apache kafka)

Tech Stack
Backend
Language: Go 1.24.1
Framework: go-restful (REST API framework)
Architecture Pattern: Layered architecture (Endpoints → Services → DAO → Models)
Databases
PostgreSQL - Primary database with gorm ORM
Redis - Caching layer (go-redis/v9)
MongoDB - Document storage (mongo-driver)
S3 - Object storage (aws-sdk-go-v2)
Authentication & Security
JWT - Token-based auth (golang-jwt/jwt)
AWS Cognito - User management integration
Observability
OpenTelemetry - Distributed tracing and metrics
Custom Logger - Cisco internal logging library
Testing
Testify - Testing framework
Go-SQLMock - Database testing
gomock - Mock generation
DevOps & Infrastructure
Docker - Containerization (Alpine-based multi-stage build)
Jenkins - CI/CD pipeline
Kubernetes - Orchestration (inferred from distributed lock patterns)
Key Dependencies
emicklei/go-restful - REST API
gorm.io/gorm - ORM for PostgreSQL
aws/aws-sdk-go-v2 - AWS services integration
go.mongodb.org/mongo-driver - MongoDB client
redis/go-redis - Redis client
Cisco internal libraries for logging, utilities, and OpenTelemetry




Inside one node we can have multiple pods(serving different micro-services) ---- good practise
POD:  one pod in kubernetes will have 3 below containers:
1) opa-istio
2) otel-coll
3) business-logic


DEPLOYMENT:

WHEN there is new microservice:
We release the opa changes (add new endpoint to the repo) this is what sets firewall to opa-istio
After github actions ci job :
which check code quality, secrets exposed , sonar cube scans, docker image vulnerability scans, builds image and pushes to aws ecr

Then comes CD job (using circleci):
we give inputs:
{
    "imageTag":"c_12hdhdwd",   // ecr's image tag which was pushed
    "ENV":{
        "port":"8084",
        "model":"value",
    },
    "New_Setup": {
        "resources": {
            "requests": {
                "cpu": "900m",      // Average baseline
                "memory": "1024Mi"  // Minimum needed
            },
            "limits": {
                "cpu": "1800m",     // 2× request for bursts
                "memory": "1536Mi"  // 1.5× request with buffer
            }
        }
    },
    "targetEnvironment":"dev"  // stage, Prod
}

// 1 cpu core = 1000 millicores
// 1gb = 1024 mb

## Vertical Pod Autoscaling (VPA)

Yes, increasing the CPU and memory of a pod in Kubernetes is considered vertical scaling, often referred to as Vertical Pod Autoscaling (VPA)

Tuning CPU:
Request = average usage
Limit = 2–3× request

Memory
Request = Limit (almost always)


## Horizontal pod auto-scaling
### How to Make It Highly Available for Thousands of Users 

CPU/memory alone won’t scale users.

You need Horizontal Pod Autoscaling (HPA)

Example:

minReplicas: 3
maxReplicas: 20
metrics:
- type: Resource
  resource:
    name: cpu
    target:
      type: Utilization
      averageUtilization: 70


Meaning:

If CPU > 70%, scale up
else scale down

Always keep minimum 3 pods



DEBUGGING MICROSERVICE in PRODUCTION:

ai-assistant -------> proxy service (taking more time)

first we go to AWS xray (now we are using splunk) and check how much time that particular api call is taking from ai-assistant
using the traces and logs(open telemetry).

Then we will get to know the complete latency and also the timeline as to how the request is evolving from one function to another inside that microservice. (using induvidual spans) 


## roles and responsiblities

![alt text](image.png)


## questions

![alt text](image-1.png)