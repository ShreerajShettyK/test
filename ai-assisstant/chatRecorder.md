┌─────────────┐         ┌──────────────┐         ┌─────────────────────┐
│   User      │────────▶│  AI Chatbot  │────────▶│  LLM (OpenAI/Claude)│
│             │         │  (Frontend)  │         │                     │
└─────────────┘         └──────────────┘         └─────────────────────┘
                               │                            │
                               │                            │
                               ▼                            ▼
                        1. Get history                 Returns answer
                        2. Build context               
                        3. Send to LLM ──────────────────┘
                        4. Record Q&A
                               │
                               ▼
                    ┌──────────────────────────┐
                    │ gen-ai-llm-chat-recorder │  ◀─── This Service
                    │  (Storage Service)       │
                    └──────────────────────────┘
                               │
                               ▼
                    ┌──────────────────────────┐
                    │ PostgreSQL + Redis + S3  │
                    └──────────────────────────┘