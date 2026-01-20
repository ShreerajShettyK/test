AI-Assistant Flow:

1) for user when he asks a question to ai-assistant:

UI frontend----> Ui backend (authorization, profile will be checked)------> chat-recorder to fetch the conversations and seed questions -----> when user enters the query ----> question will be taken from ui to ui 's backend ----> UI backend ------> Secure shield (USES gemini model) (blocks harmfull question) (true or false) ------->
if attack detected is false 
-----> go to learning-assistant service
                                        ---->
                                            citation (sources related to the questions asked like videos, podcasts, etc from cisco db) ---------------> embeddings of all cisco data

                                        ----->
                                             related questions---- fecthes all the questions related to the question the user asked

                                        ------>
                                             topic mapper---- creates graph which gives a visual representation of all the topics related to the one which was aked

                                        ------>chat recorder--- to fetch chats and convos (for context)


                                        ------> cache service     


                                        -------> Proxy service - talks to llms azure openai, aws bedrock, circuit, apim, gemini


                                        -------> Embedding service (used by citation and toppic mapper) uses half vctor (not full vector)



2) for users who asks questions related to training questions

 Training recommender (for training related questions)

3) Testing

 gen-ai-hammer -------- used for testing perfomrance of ai powered applications using vegeta and hey packages                                       

4) Used by analytics team so that they get to know 

 gen-ai-reports-service --- used by analytics
 takes all the input qa pairs and in batches send it to axure model and groups it to categories defined in the prompt and excel will be uploaded to s3




EXAM GENERATOR (new feature) FLOW:


1) for the admin (subject matter expert)---------------------------------kafka
ui backend calls for admin ops below//

                task.preflight passed                           task.questionsgenerated
 exam assistant ----------------------> question generator -----------------------------> secure-shield 
 
 task.securitypassed                                           task.readyforReview
 -----------------------------> sme_router(stores to db) -------------------------------> exam assistant


 then after all the questions are ready and generated. The exam assistant will call the sme_router for displaying the questions. Till then it will be in progress state.