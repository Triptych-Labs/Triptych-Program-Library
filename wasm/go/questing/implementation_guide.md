# WASM Implementation for Frontend Guide

The following are coded: `order-priority` and must be executed in order and in priority.


## Landing Page
### 1-a) Fetch Quests
use `get_quests`.

### 1-b) Fetch Quests Proposals
then use `get_quests_proposals`.


## Staking Page (and/or)
### 2-a) Create Quest Record (if needed, on load)
use `select_quest`. this will create a record of all the proposals for this wallet and quest.

### 3-a) Initialize (new) Quest Proposal && Enter/deposit stakes (on finish)
use `new_quest_proposal`. this will create a new proposal and deposits transaction.


## Splash Page
### 4-a) Start Quest

