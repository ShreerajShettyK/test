┌─────────┐
│  User   │
└────┬────┘
     │ "What is Python?"
     ▼
┌──────────────────────┐
│   AI Chatbot         │
│   (Frontend)         │
└──────────────────────┘
     │
     ▼
┌──────────────────────────────────────────────────────────┐
│  1. Check Cache Service (NEW)                            │
│     POST /api/cache/query                                │
│     { "question": "What is Python?", "context": "Last 10 chats summarized(got from chat-recorder)" }  │
└──────────────────────────────────────────────────────────┘
     │
     ├─── CACHE HIT ────────────┐
     │    (Similarity > 95%)     │
     │                           ▼
     │                    Return cached answer
     │                    (0.1s response time)
     │                    from_cache: true
     │
     └─── CACHE MISS ───────────┐
          (No similar question)  │
                                ▼
                    ┌────────────────────────┐
                    │  Call Proxy Service    │
                    │  (OpenAI/Claude)       │
                    │  (2-5s response time)  │
                    └────────────────────────┘
                                │
                                ▼
                    ┌────────────────────────┐
                    │  Cache the Q&A pair    │
                    │  POST /api/cache/store │
                    └────────────────────────┘
                                │
                                ▼
                    ┌───────────────────────────────┐
                    │  Record to chat-recorder      │
                    │  with from_cache=false        │
                    └───────────────────────────────┘




->Vector Database for Semantic Search:

Using pgvector (PostgreSQL extension):

Step 1: Embedding Generation
When a question enters the system, it goes through an embedding model (like OpenAI's text-embedding-ada-002 or open-source models):


┌───────────────────────────────────────────────────────┐
│ Input to Embedding Model:                            │
├───────────────────────────────────────────────────────┤
│                                                       │
│ Question: "How do I install it?"                     │
│                                                       │
│ Context Summary(which we got from chat-recorder)      │
│ - User previously asked about Python                  │
│ - Discussion topic: programming languages             │
│ - Agent type: coding-assistant                        │
│                                                       │
│ Combined Text for Embedding:                          │
│ "Context: Python programming language discussion.     │
│  Question: How do I install it?"                      │
│                                                       │
│        ↓                                              │
│     output                                            │
│ Vector: [0.82, 0.31, -0.15, 0.44, ...]                │
│ (This vector now encodes both question AND context)   │
└───────────────────────────────────────────────────────┘


            Input Text:
                 ↓
        [Embedding Model]
                 ↓
Output Vector: [0.023, -0.891, 0.445, 0.112, -0.334, 0.667, ...]
               (1536 numbers representing semantic meaning)