# IBC
## CSE528 Introduction to Blockchain and Cryptocurrencies
### Topic
#### Using Hyperledger Fabric to manage Human Aging Genomics Data

##### Introduction
To develop a robust method for storing, sharing and updating human aging genomics data among different entities like pharm companies, academics, researchers and healthcare professionals so that the data can be selectively shared while maintaining privacy. Users own the data and can decide to share their genomics data anonymously, researchers get credit for their contributions, pharmaceutical companies get paid for their research and development. 

To build a proof of concept for the blockchain based human aging genomics data storage.
The Human Ageing Genomic Resources (HAGR) is a collection of databases and tools designed to help researchers study the genetics of human ageing using modern approaches such as functional genomics, network analyses, systems biology and evolutionary analyses.

### Implementation
Programming plan to implement the proposed project using open source hyperledger fabric.

#### March Objective

##### Week 1 (Ideation)
* Project ideation and conceptualization of human aging genomic data privacy.
* Idea validation and read case studies on blockchain based genomics startups
* Evaluate existing blockchain business models in the genomics space.
 
##### Week 2 (Fabric Fundamentals)
Learn hyperledger fabric framework, tools and component design through tutorials. 
Getting familiar with Hyperledger Fabric platform and terminologies.
Installation of prerequisites and Fabric SDK to work on blockchain projects.
Run any demo project from fabric samples for understanding.

##### Week 3 (Project framework)
Define Fabric chaincode lifecycle for interaction between different participants.
Using Private Data in Fabric and making collections for data read write accessibility.
Define assets, transactions, events and understand business logic of the network.
Defining endorsement policies for different members in the organisation prior to establishing a channel in the network.

##### Week 4 (Database creation)
Using CouchDB database to store human aging genomics data.
Learn about couchDB from video CouchDB for Fabric Developers and other sources.
Define onchain metadata and hash to be stored in the transaction on blockchain and off chain immutable large genomic data. 
Preparing sideDB database used by Hyperledger fabric for storing genomic data.

#### April Objective

##### Week 1 (Fabric Network)
Create different types of participants at individual, organisation and system level.
Assign participants roles and responsibilities like who can access data, create a channel, take part in consensus mechanisms. 
Design structure of assets and their ownerships and how transactions are verified.
Design selective data sharing based on transaction and sideDB database.

##### Week 2 (Develop Smart Contract)
To design and implement smart contract transactions and ledger data structures
Implement access control, code smart contract functions to read and modify assets.

##### Week 3 (Fabric Security)
Design and implement collection, state endorsement policies and ordering service.
Define Fabric CA (certificate authority) and membership service provider.
Use SDK to access fabric networks, submit transactions and query ledger and listen to the response from the network.

##### Week 4 (Deploy Fabric Network)
Run blockchain application, test and debug any errors.
Check functionality of the code and modify as needed.
Make sure the code is written with proper comments.

#### References
1. Hyperledger Fabric A Blockchain Platform for the Enterprise
2. SideDB pptx file Privacy Enabled Ledger
3. Tacutu, R., et al. (2018) "Human Ageing Genomic Resources: new and updated databases." Nucleic Acids Research 46(D1):D1083-D1090
4. Gürsoy, G., Brannon, C.M. & Gerstein, M. Using Ethereum blockchain to store and query pharmacogenomics data via smart contracts. BMC Med Genomics 13, 74 (2020). https://doi.org/10.1186/s12920-020-00732-x
5. Zero Knowledge Proof https://en.wikipedia.org/wiki/Zero-knowledge_proof
