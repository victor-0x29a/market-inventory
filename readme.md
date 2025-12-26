# Requirements

- C Compiler (for run specs using that's use SQLITE)
- Golang
- Docker, to run locally

# Market Inventory

API built in Go for market inventory management, allowing product control and damage logging.

## Overview

The system allows managing products and registering damage logs associated with inventory items. It also provides access to predefined damage reasons. Listing endpoints are paginated and currently do not support filters.

## API Routes

### Products

<pre class="overflow-visible! px-0!" data-start="439" data-end="579"><div class="contain-inline-size rounded-2xl corner-superellipse/1.1 relative bg-token-sidebar-surface-primary"><div class="overflow-y-auto p-4" dir="ltr"><code class="whitespace-pre!"><span><span>POST   /v1/products
GET    /v1/products
GET    /v1/products/:productId
PATCH  /v1/products/:productId
DELETE /v1/products/:productId
</span></span></code></div></div></pre>

### Damage LogsPOST /v1/damage-log

### Damage Reasons

<pre class="overflow-visible! px-0!" data-start="667" data-end="701"><div class="contain-inline-size rounded-2xl corner-superellipse/1.1 relative bg-token-sidebar-surface-primary"><div class="overflow-y-auto p-4" dir="ltr"><code class="whitespace-pre!"><span><span>GET /v1/damage-log/reasons
</span></span></code></div></div></pre>

---
