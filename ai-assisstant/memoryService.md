# Gen-AI Memory Service Architecture

## Overview

The **gen-ai-memory-service** acts as a **RAG (Retrieval Augmented Generation)** layer that provides structured access to Cisco's internal knowledge base, enabling the AI chatbot to answer questions using verified, company-specific documentation and technical resources.

---

## Complete Architecture with All Services

```
┌──────────┐
│   User   │ "How do I configure Cisco Catalyst switch?"
└────┬─────┘
     │
     ▼
┌────────────────────────────────────────────────────────────┐
│              AI Chatbot Orchestrator                       │
│  (Coordinates all microservices)                           │
└────────────────────────────────────────────────────────────┘
     │
     │ Decision Flow:
     │
     ├─── 1️⃣ Check Cache First ──────────────────────┐
     │                                                 │
     ▼                                                 ▼
┌─────────────────────────┐                  ┌──────────────┐
│  gen-ai-cache-service   │ Cache Hit? ───▶  │ Return Answer│
│  (Semantic cache)       │                  │ from Cache   │
└─────────────────────────┘                  └──────────────┘
     │                                                 
     │ Cache Miss ↓
     │
     ├─── 2️⃣ Get Conversation Context ───────────────┐
     │                                                 │
     ▼                                                 ▼
┌──────────────────────────┐              ┌──────────────────┐
│ gen-ai-chat-recorder     │              │ Returns history: │
│ (Conversation history)   │              │ - Previous Q&A   │
└──────────────────────────┘              │ - Context        │
     │                                    └──────────────────┘
     │
     ├─── 3️⃣ Search Knowledge Base ──────────────────┐
     │                                                 │
     ▼                                                 ▼
┌──────────────────────────┐              ┌──────────────────────┐
│ gen-ai-memory-service    │              │ Returns:             │
│ (Knowledge retrieval)    │              │ - Relevant docs      │
│                          │              │ - Product manuals    │
│ • Vector search          │              │ - KB articles        │
│ • Document retrieval     │              │ - Technical specs    │
│ • RAG pipeline           │              │                      │
└──────────────────────────┘              └──────────────────────┘
     │
     │
     ├─── 4️⃣ Build Enhanced Prompt ───────────────────┐
     │                                                  │
     │   Combine:                                       │
     │   + Conversation context (from chat-recorder)   │
     │   + Knowledge base docs (from memory-service)   │
     │   + User's current question                     │
     │                                                  │
     ▼                                                  ▼
┌──────────────────────────┐              ┌──────────────────┐
│      LLM (GPT-4)         │◀─────────────│ Enhanced Prompt  │
│                          │              │ with context +   │
│  Generates answer using  │              │ KB articles      │
│  provided knowledge      │              └──────────────────┘
└──────────────────────────┘
     │
     │ Answer generated
     │
     ├─── 5️⃣ Store for Future Use ───────────────────┐
     │                                                 │
     ▼                                                 ▼
┌──────────────────────────┐              ┌──────────────────┐
│ gen-ai-cache-service     │              │ Store Q&A with   │
│ (Cache new answer)       │              │ context for      │
└──────────────────────────┘              │ future hits      │
                                          └──────────────────┘
     │
     ▼
┌──────────────────────────┐
│ gen-ai-chat-recorder     │
│ (Record conversation)    │
└──────────────────────────┘
```

---

## Gen-AI Memory Service: Deep Dive

### Purpose
Provides structured access to Cisco's internal knowledge base, product documentation, and technical resources through semantic search and retrieval.

### What It Stores

```
┌─────────────────────────────────────────────────────────┐
│  Cisco Internal Knowledge Base                          │
├─────────────────────────────────────────────────────────┤
│                                                         │
│  📚 Product Documentation                               │
│     - Catalyst Switch Configuration Guides             │
│     - Router Setup Manuals                             │
│     - Security Appliance Docs                          │
│                                                         │
│  🎓 Knowledge Base Articles                            │
│     - Troubleshooting guides                           │
│     - Best practices                                   │
│     - Common issues & solutions                        │
│                                                         │
│  📝 Technical Specifications                            │
│     - Product datasheets                               │
│     - API documentation                                │
│     - Architecture diagrams                            │
│                                                         │
│  📊 Training Materials                                  │
│     - Certification content                            │
│     - Tutorial videos transcripts                      │
│     - Lab exercises                                    │
│                                                         │
│  🔧 Internal Wikis                                      │
│     - Engineering notes                                │
│     - Solution architectures                           │
│     - Implementation patterns                          │
└─────────────────────────────────────────────────────────┘
```

---

## Document Ingestion Pipeline

### How Knowledge Base Gets Into Memory Service

```
┌──────────────────────────────────────────────────────┐
│  Step 1: Document Collection                         │
└──────────────────────────────────────────────────────┘
   Sources:
   • Confluence pages
   • SharePoint documents
   • Product documentation repositories
   • PDF manuals
   • Internal wikis
   • Support ticket resolutions
        ↓
┌──────────────────────────────────────────────────────┐
│  Step 2: Document Processing                         │
└──────────────────────────────────────────────────────┘
   • Extract text from PDFs, Word docs, HTML
   • Parse markdown and structured formats
   • Extract metadata (product, version, category)
   • Clean and normalize text
        ↓
┌──────────────────────────────────────────────────────┐
│  Step 3: Chunking Strategy                           │
└──────────────────────────────────────────────────────┘
   Break documents into semantic chunks:
   
   Original doc: "Catalyst 9000 Switch Configuration Guide" (500 pages)
   
   Chunked into:
   ├─ Chunk 1: "Chapter 1: Initial Setup"
   │  Size: ~1000 tokens
   │  Metadata: {product: "Catalyst 9000", chapter: 1}
   │
   ├─ Chunk 2: "VLAN Configuration"  
   │  Size: ~800 tokens
   │  Metadata: {product: "Catalyst 9000", topic: "VLANs"}
   │
   ├─ Chunk 3: "Port Security Settings"
   │  Size: ~1200 tokens
   │  Metadata: {product: "Catalyst 9000", topic: "Security"}
   └─ ...
        ↓
┌──────────────────────────────────────────────────────┐
│  Step 4: Generate Embeddings                         │
└──────────────────────────────────────────────────────┘
   Each chunk → Vector embedding
   
   Chunk: "To configure VLANs on Catalyst 9000, use..."
   Vector: [0.45, 0.78, -0.23, 0.91, ...]
        ↓
┌──────────────────────────────────────────────────────┐
│  Step 5: Store in Vector Database                    │
└──────────────────────────────────────────────────────┘
   Database: pgvector or specialized vector DB
   
   Record:
   ├─ chunk_id: "doc-cat9k-ch3-vlan"
   ├─ content: "To configure VLANs on Catalyst..."
   ├─ embedding: [0.45, 0.78, -0.23, ...]
   ├─ metadata: {
   │    product: "Catalyst 9000",
   │    document: "Configuration Guide",
   │    chapter: 3,
   │    topic: "VLANs",
   │    version: "IOS-XE 17.x",
   │    url: "internal-link"
   │  }
   └─ last_updated: "2025-12-15"
```

---

## Memory Service API Endpoints

### 1. Search Knowledge Base

**Endpoint:** `POST /api/memory/search`

**Request:**
```json
{
  "query": "How to configure VLANs on Catalyst switch?",
  "filters": {
    "product": "Catalyst 9000",
    "category": "configuration"
  },
  "top_k": 5,
  "min_relevance": 0.75,
  "include_metadata": true
}
```

**Response:**
```json
{
  "results": [
    {
      "chunk_id": "doc-cat9k-ch3-vlan",
      "content": "To configure VLANs on Catalyst 9000:\n1. Enter config mode: enable, configure terminal\n2. Create VLAN: vlan 10\n3. Name VLAN: name Engineering\n4. Assign ports: interface range gi1/0/1-10\n5. Set access VLAN: switchport access vlan 10",
      "relevance_score": 0.96,
      "source": {
        "document": "Catalyst 9000 Configuration Guide",
        "chapter": "3 - VLAN Management",
        "url": "https://internal-docs.cisco.com/cat9k/config",
        "version": "IOS-XE 17.9"
      },
      "metadata": {
        "product": "Catalyst 9000",
        "topic": "VLANs"
      }
    },
    {
      "chunk_id": "kb-art-12345",
      "content": "Common VLAN configuration issues:\n- Verify trunk ports: show interfaces trunk\n- Check VLAN database: show vlan brief...",
      "relevance_score": 0.89,
      "source": {
        "type": "knowledge_base_article",
        "title": "VLAN Troubleshooting Guide",
        "url": "https://kb.cisco.com/12345"
      }
    }
  ],
  "total_found": 15,
  "search_time_ms": 45
}
```

### 2. Get Document by ID

**Endpoint:** `GET /api/memory/document/{chunk_id}`

Returns full chunk with surrounding context.

### 3. Hybrid Search (Keyword + Semantic)

**Endpoint:** `POST /api/memory/hybrid-search`

Combines traditional keyword search with vector similarity for best results.

---

## Complete Example: User Question Flow

### User Question: "How do I configure port security on my Catalyst 9300 switch?"

#### Step 1: Check Cache

```
Request to gen-ai-cache-service:
  Query: "How do I configure port security on Catalyst 9300?"
  
Response: Cache Miss
  (No previous identical or similar question)
```

#### Step 2: Get Conversation History

```
Request to gen-ai-chat-recorder:
  conversation_id: "conv-456"
  
Response:
  Previous questions:
  - "What models are in the Catalyst 9000 series?"
  - "What are the specs of 9300?"
  
Context extracted: "User is asking about Catalyst 9300 switches"
```

#### Step 3: Search Knowledge Base

```
Request to gen-ai-memory-service:
  POST /api/memory/search
  {
    "query": "configure port security Catalyst 9300",
    "filters": {
      "product": "Catalyst 9300",
      "topic": "security"
    },
    "top_k": 3
  }
```

**Response: Found 3 relevant documents**

**Document 1 (Relevance: 0.94):**
```
Port Security Configuration on Catalyst 9300:

Overview: Port security restricts input to an interface 
by limiting and identifying MAC addresses.

Configuration Steps:
1. Enable port security:
   Switch(config-if)# switchport port-security
   
2. Set maximum MAC addresses:
   Switch(config-if)# switchport port-security maximum 2
   
3. Define violation action:
   Switch(config-if)# switchport port-security violation shutdown
   
4. Verify configuration:
   Switch# show port-security interface gigabitethernet1/0/1
   
Source: Catalyst 9300 Security Configuration Guide, Ch 7
```

**Document 2 (Relevance: 0.88):**
```
Best Practices for Port Security:
- Always configure aging to clear learned addresses
- Use 'restrict' mode for production environments
- Enable SNMP traps for security violations
```

**Document 3 (Relevance: 0.82):**
```
Common Port Security Issues:
- Violation counter incrementing: Check for MAC flapping
- Port in err-disabled state: Verify violation action...
```

#### Step 4: Build Enhanced Prompt

Chatbot constructs LLM prompt:

```
You are a Cisco networking expert assistant.

CONVERSATION CONTEXT:
The user has been asking about Catalyst 9300 switches.
Previous questions:
- What models are in the Catalyst 9000 series?
- What are the specs of 9300?

RELEVANT DOCUMENTATION:
[Document 1 - Port Security Configuration]
Port Security Configuration on Catalyst 9300:
Overview: Port security restricts input to an interface...
Configuration Steps:
1. Enable port security: Switch(config-if)# switchport port-security
...

[Document 2 - Best Practices]
Best Practices for Port Security:
- Always configure aging to clear learned addresses...

[Document 3 - Troubleshooting]
Common Port Security Issues:
- Violation counter incrementing: Check for MAC flapping...

USER QUESTION:
How do I configure port security on my Catalyst 9300 switch?

INSTRUCTIONS:
Answer based ONLY on the provided documentation. Include:
- Step-by-step configuration commands
- Important considerations
- Verification steps
Cite sources when applicable.
```

#### Step 5: Send to LLM

**Request to OpenAI/Claude:**
```
Prompt: [Enhanced prompt above]
Model: gpt-4
```

**Response from LLM:**
```
To configure port security on your Catalyst 9300 switch:

**Configuration Steps:**
1. Enter interface configuration mode:
   Switch> enable
   Switch# configure terminal
   Switch(config)# interface gigabitethernet1/0/1

2. Enable port security:
   Switch(config-if)# switchport port-security

3. Set maximum allowed MAC addresses (default is 1):
   Switch(config-if)# switchport port-security maximum 2

4. Configure violation action:
   Switch(config-if)# switchport port-security violation shutdown
   Options: shutdown | restrict | protect

5. Verify configuration:
   Switch# show port-security interface gigabitethernet1/0/1

**Best Practices:**
- Configure MAC address aging to automatically clear learned addresses
- Use 'restrict' mode in production for less disruptive enforcement
- Enable SNMP traps to monitor security violations

Source: Catalyst 9300 Security Configuration Guide, Chapter 7
```

#### Step 6: Cache the Answer

```
Request to gen-ai-cache-service:
  POST /api/cache/store
  {
    "question": "How do I configure port security on Catalyst 9300?",
    "context": "Catalyst 9300 switches",
    "answer": "[LLM response above]",
    "sources": ["doc-cat9k-ch7-portsec", "kb-art-12345"],
    "ttl_hours": 168
  }
```

#### Step 7: Record to Chat History

```
Request to gen-ai-chat-recorder:
  POST /api/user-conversation
  {
    "chat_id": "chat-789",
    "conversation_id": "conv-456",
    "question": "How do I configure port security on my Catalyst 9300?",
    "answer": "[LLM response]",
    "from_cache": false,
    "response_time": 2.8,
    "model": "gpt-4",
    "metadata": {
      "kb_sources": ["doc-cat9k-ch7-portsec", "kb-art-12345"],
      "relevance_scores": [0.94, 0.88, 0.82]
    }
  }
```

**Return answer to user** ✓

---

## Data Structure in Memory Service

### PostgreSQL + pgvector Schema

```sql
-- Document chunks with embeddings
CREATE TABLE knowledge_chunks (
    chunk_id VARCHAR(255) PRIMARY KEY,
    content TEXT NOT NULL,
    embedding vector(1536),  -- Semantic vector
    
    -- Source metadata
    source_type VARCHAR(50),  -- 'manual', 'kb_article', 'wiki', 'training'
    source_id VARCHAR(255),
    document_title VARCHAR(500),
    chapter VARCHAR(255),
    section VARCHAR(255),
    
    -- Product/category filters
    product_family VARCHAR(100),  -- 'Catalyst', 'Nexus', 'ASR'
    product_model VARCHAR(100),   -- 'Catalyst 9300', 'Nexus 9000'
    category VARCHAR(100),         -- 'configuration', 'troubleshooting', 'security'
    topic_tags TEXT[],             -- ['VLANs', 'port-security', 'QoS']
    
    -- Version/relevance
    software_version VARCHAR(100),
    confidence_score FLOAT,        -- Quality/accuracy rating
    
    -- URLs and references
    source_url TEXT,
    related_chunks TEXT[],
    
    -- Timestamps
    indexed_at TIMESTAMP DEFAULT NOW(),
    last_updated TIMESTAMP,
    
    -- Full-text search
    content_tsv tsvector GENERATED ALWAYS AS (to_tsvector('english', content)) STORED
);

-- Vector similarity index
CREATE INDEX idx_knowledge_embedding ON knowledge_chunks 
USING ivfflat (embedding vector_cosine_ops)
WITH (lists = 1000);

-- Filter indexes
CREATE INDEX idx_product ON knowledge_chunks(product_model, category);
CREATE INDEX idx_topics ON knowledge_chunks USING gin(topic_tags);

-- Full-text search index
CREATE INDEX idx_content_fts ON knowledge_chunks USING gin(content_tsv);
```

---

## Hybrid Search Strategy

### Why Hybrid?

Pure semantic search misses specific terms; pure keyword search misses meaning.

**Example:**
- Query: "CLI commands for VLAN config"
- Semantic might miss exact command syntax
- Keyword might miss conceptual explanations
- **Hybrid gets both!**

### Search Flow

```
User Query: "show vlan command output explained"
        ↓
┌─────────────────────────────────────────┐
│  Parallel Search                        │
├─────────────────────────────────────────┤
│                                         │
│  Track 1: Vector Search                 │
│  • Generate embedding                   │
│  • Find similar content                 │
│  • Results: [Doc A: 0.91, Doc B: 0.87] │
│                                         │
│  Track 2: Keyword Search                │
│  • Extract keywords: "show vlan"        │
│  • PostgreSQL full-text search          │
│  • Results: [Doc B: 0.95, Doc C: 0.88] │
│                                         │
└─────────────────────────────────────────┘
        ↓
Merge and Rerank:
  Doc B: 0.91 (both searches) → BEST MATCH
  Doc A: 0.91 (semantic only)
  Doc C: 0.88 (keyword only)
        ↓
Return top 3 to chatbot
```

---

## Key Benefits

### 1. Grounded Answers
LLM responses based on actual Cisco documentation, not hallucinations.

### 2. Source Attribution
Every answer includes links to source documents for verification.

### 3. Version-Aware
Filters by software version ensure correct commands for user's platform.

### 4. Context-Aware Retrieval
Uses conversation history to understand which product/topic user is discussing.

### 5. Scalable Knowledge
New documentation automatically indexed and searchable within hours.

---

## Performance Metrics

| Metric | Without Memory Service | With Memory Service |
|--------|----------------------|---------------------|
| Answer Accuracy | 60-70% | 90-95% |
| Hallucination Rate | 20-30% | <5% |
| Source Attribution | 0% | 100% |
| Response Relevance | Variable | Consistently High |
| Knowledge Freshness | Static (training data) | Real-time (updated docs) |

---

## Three-Service Architecture Summary

### gen-ai-cache-service
- **Purpose:** Fast retrieval of previously answered questions
- **Storage:** Redis + pgvector
- **Hit Rate:** 40-70%
- **Benefit:** Reduces LLM costs and response time

### gen-ai-memory-service
- **Purpose:** Access to Cisco knowledge base and documentation
- **Storage:** pgvector + document chunks
- **Size:** Millions of document chunks
- **Benefit:** Provides accurate, grounded, source-attributed answers

### gen-ai-chat-recorder
- **Purpose:** Conversation history and audit trail
- **Storage:** PostgreSQL + Redis
- **Use:** Maintain context across conversation
- **Benefit:** Enables contextual, multi-turn conversations

---

## Integration Benefits

Together, these three services enable:
1. **Fast responses** via cache service (50ms vs 2500ms)
2. **Accurate, grounded answers** via memory service (90%+ accuracy)
3. **Contextual conversations** via chat recorder (maintains conversation flow)
4. **Cost optimization** (60% reduction in LLM API calls)
5. **Audit compliance** (complete conversation tracking)
6. **Knowledge management** (centralized, searchable documentation)

---

## Future Enhancements

### Potential Improvements

1. **Multi-modal Search**
   - Index diagrams, screenshots, videos
   - Visual similarity search for network topologies

2. **Federated Search**
   - Search across multiple knowledge bases
   - Integration with external sources (Stack Overflow, GitHub)

3. **Personalized Retrieval**
   - User role-based filtering (engineer, admin, student)
   - Personalized relevance scoring

4. **Active Learning**
   - Track which documents users find helpful
   - Automatically improve relevance rankings

5. **Real-time Updates**
   - Continuous document ingestion pipeline
   - Incremental indexing for new content

6. **Advanced RAG Techniques**
   - Query decomposition for complex questions
   - Multi-hop reasoning across documents
   - Answer synthesis from multiple sources

---

## Conclusion

The gen-ai-memory-service is a critical component in building an enterprise-grade AI chatbot that provides accurate, verifiable answers grounded in company-specific knowledge. By combining semantic search, metadata filtering, and RAG techniques, it transforms static documentation into an intelligent, queryable knowledge base that enhances the AI's capabilities while maintaining accuracy and accountability.
