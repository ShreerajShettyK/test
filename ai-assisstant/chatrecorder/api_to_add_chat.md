curl --location 'http://localhost:8086/gen-ai-llm-chat-recorder/api/chat' \
--header 'Authorization: Bearer eyJraWQiOiJOaDE0VENXc29FZWVNb2hPMW5nNHhkbnp3UURmaG1hN0RBUmRUcDJtVlk0PSIsImFsZyI6IlJTMjU2In0.eyJzdWIiOiIzdGkxaTdnNm5hajFkN2xzOGxjOHZqMjZwZCIsInRva2VuX3VzZSI6ImFjY2VzcyIsInNjb3BlIjoibGNwX3NlcnZpY2VcL2NvbnRlbnQtYXV0aHoucmVhZCBsY3Bfc2VydmljZV9nZW5fYWlcL2dlbi1haS1sbG0tcHJveHkucmVhZCBsY3Bfc2VydmljZV9nZW5fYWlcL2dlbi1haS1sbG0tcHJveHkud3JpdGUgbGNwX3NlcnZpY2VfZ2VuX2FpXC9nZW4tYWktcmF0ZS1saW1pdC1zZXJ2aWNlLnJlYWQgbGNwX21pY3Jvc2VydmljZVwvZ2VuLWFpLWNhY2hlLXNlcnZpY2UucmVhZCBsY3BfbWljcm9zZXJ2aWNlXC9nZW4tYWktbGxtLWNsaWVudC5yZWFkIGxjcF9zZXJ2aWNlXC9jb250ZW50LWNhdGFsb2cud3JpdGUgbGNwX3NlcnZpY2VcL3Byb2ZpbGUtc2VydmljZS5yZWFkIGxjcF9taWNyb3NlcnZpY2VcL2dlbi1haS1sbG0tY2hhdC1yZWNvcmRlci1zZXJ2aWNlLnJlYWQgbGNwX21pY3Jvc2VydmljZVwvZ2VuLWFpLWNhY2hlLXNlcnZpY2Uud3JpdGUgbGNwX3NlcnZpY2VfZ2VuX2FpXC9nZW4tYWktcHJvbXB0LXBhcnNlci53cml0ZSBsY3BfbWljcm9zZXJ2aWNlXC9nZW4tYWktbGxtLWNoYXQtcmVjb3JkZXItc2VydmljZS53cml0ZSBsY3Bfc2VydmljZVwvY29udGVudC1hdXRoei53cml0ZSBsY3Bfc2VydmljZV9nZW5fYWlcL2dlbi1haS1wcm9tcHQtcGFyc2VyLnJlYWQgbGNwX21pY3Jvc2VydmljZVwvZ2VuLWFpLWxsbS1jbGllbnQud3JpdGUgbGNwX3NlcnZpY2VfZ2VuX2FpXC9nZW4tYWktcmF0ZS1saW1pdC1zZXJ2aWNlLndyaXRlIiwiYXV0aF90aW1lIjoxNzQ0NzE0NTk4LCJpc3MiOiJodHRwczpcL1wvY29nbml0by1pZHAudXMtZWFzdC0xLmFtYXpvbmF3cy5jb21cL3VzLWVhc3QtMV92b0EwR3BuY1MiLCJleHAiOjE3NDQ3NTc3OTgsImlhdCI6MTc0NDcxNDU5OCwidmVyc2lvbiI6MiwianRpIjoiNmU0Mzg3OTUtNWQ5Mi00ZmQ5LTk0MmYtOWEwYWE3ZTRkM2ZjIiwiY2xpZW50X2lkIjoiM3RpMWk3ZzZuYWoxZDdsczhsYzh2ajI2cGQifQ.kIhhy6yZ98deofEW_P3sfwnInSUi38W0B0C0I3ApKGItsSPcvCq3s3UjUED0njeGYQZ84QunLuiEdLhvI6Bso1tU2mQ4KtUlzpGC6PzkBRId6NBBXHQlWUCFislN0bOd4rdwQoLJA-nW5hr_cEC0ko7PY208VO8HdXeEoFpAex0AgpwZCXLFHx8aNuV-j4feJfZaZ0r0FclHjs6QtSG2FeRxDsW1XsOOVV5Gp058x_U-akxyJDLDuHIM3h7c6TDNGLuGMzVtNBWPFqEhcOf8vqt1bnlstJji0v8tnU1ULIPdGrhi_fdKOy1xeXH3oAc214xUPZKrjRVqOy2zRVk85g' \
--header 'Content-Type: application/json' \
--data '{
    "chat_id": "gde92bf2-5d96-4e74-a06a-f0616cd9e721",
    "conversation_id": "g0bd2e29-6086-405f-bb14-af42a3467c25",
    "username": "Shreeraj",
    "question": "Explain ipv186?",
    "answer": "netwroking is the most recent version of the Internet Protocol...",
    "comments": "feedback",
    "is_seed_prompt": true,
    "model": "bridgeit",
    "model_type": "albert",
    "agentType": "learn",
    "response_time": 1.23,
    "delivery_mode": 2,
    "feature": "expertContent",
    "feedback": 1,
    "rating": 3.5,
    "routing": "direct",
    "from_cache": false
}



# Payload:
## first chat endpoint

// {
//   "chat_id": "chatcmpl-8277a7a181654f42ad6a6b435b9b2b69",
//   "username": "CiscoSAML_saikraja@cisco.com",
//   "question": "What career paths open up after obtaining a Cisco Certified Network Professional (CCNP) certification?",
//   "answer": "Obtaining a Cisco Certified Network Professional (CCNP) certification can open doors to a variety of advanced",
//   "agentType": "study",
//   "agentUUID": "0063003300520031005a0048006b003d",
//   "response_time": 4.73,
//   "delivery_mode": 2,
//   "feature": "expertContent"
// } 


## follow-up:

// {
//   "conversation_id": "d2928370-00a2-45d5-9db1-38f964d8f871",
//   "chat_id": "chatcmpl-1c2df77a702449f59976420401590bd1",
//   "username": "CiscoSAML_saikraja@cisco.com",
//   "question": "What career benefits can you gain by earning a Cisco certification?",      
//   "answer": "Earning a Cisco certification offers significant career benefits, enhancing both employment opportunities and"
//   "agentType": "study",
//   "agentUUID": "0063003300520031005a0048006b003d",
//   "response_time": 4.31,
//   "delivery_mode": 2,
//   "feature": "expertContent"
// } 



























'