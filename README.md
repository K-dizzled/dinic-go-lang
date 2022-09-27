# Interns from MCS-SPBU

<table>
  <tr><td><b>Scope</b></td><td>Global</td></tr>
  <tr><td><b>Status</b></td><td>Approved</td></tr>
  <tr><td><b>Created</b></td><td>2022-08-30</td></tr>
  <tr><td><b>Updated</b></td><td>2022-09-20</td></tr>
</table>

<table>
  <tr><td><b>Responsible</b></td><td>Oleg Tsarev &ltot@project-plato.com&gt</td></tr>
  <tr><td><b>Accountable</b></td><td>Oleg Tsarev &ltot@project-plato.com&gt</td></tr>
  <tr><td><b>Consulted</b></td><td>Konstantin Mashukov &ltkm@project-plato.com&gt</td></tr>
  <tr><td><b>Informed</b></td><td>Everybody</td></tr>
</table>

- [Interns from MCS-SPBU](#interns-from-mcs-spbu)
  - [Summary](#summary)
  - [Proposed Solution](#proposed-solution)
    - [General](#general)
    - [Timeline](#timeline)
      - [1. Introduction](#1-introduction)
      - [2. Evaluation](#2-evaluation)
      - [3. Learning package](#3-learning-package)
      - [4. Onboarding (checkpoint 1)](#4-onboarding-checkpoint-1)
      - [5. Dealing with JSON-LD documents (checkpoint 2)](#5-dealing-with-json-ld-documents-checkpoint-2)
      - [6. Loading JSON-LD documents to Neptune (checkpoint 3)](#6-loading-json-ld-documents-to-neptune-checkpoint-3)
      - [7. Querying documents in Neptune (checkpoint 4)](#7-querying-documents-in-neptune-checkpoint-4)
      - [8. Updating graph representations by new documents (checkpoint 5)](#8-updating-graph-representations-by-new-documents-checkpoint-5)
      - [9. Summarizing the result (checkpoint 6)](#9-summarizing-the-result-checkpoint-6)
      - [10. Applying results to work](#10-applying-results-to-work)
  - [Follow-up Tasks](#follow-up-tasks)

## Summary

[Dmitry Shalymov](https://math-cs.spbu.ru/en/people/dmitry-s-shalymov/) - the Head of the [Modern Software Engineering Program](https://bsse.compscicenter.ru/) of the 
[Mathematics and Computer Science Department](https://math-cs.spbu.ru) of Saint-Peterburg University proposed an opportunity to get interns to solve some practical industry tasks (more [information](https://math-cs.spbu.ru/projects-studentsmcs/) about these projects - Russian).

These students:

- Are winners of math olympiads (e.g. [INTERNATIONAL OLYMPIAD IN INFORMATICS](https://ioinformatics.org), 
[International Mathematical Olympiad](https://www.imo-official.org/?language=en), All-Russian Olympiads in Mathematics and Informatics).
- Know English at the writing and speaking level.
- Have an excellent mathematical background (for instance, in discrete math - i.e., graphs). The founder and scientific director of the faculty are Konstantin Smirnov, winner of the Fields Prize.
Therefore, mathematics at the faculty is now one of the best in the Russian Federation.
- Are highly motivated to do some practical industry work to receive real-life experience.
- This work is free of charge.

To move forward with this initiative, we should write a proposal describing how we will proceed.
It could be an excellent way to do a bunch of R&D related to Neptune or TigerGraph databases. 
We will NOT receive the excellent quality source code, but we will receive a dirty but workable solution as a practical starting point to design the proper solution on our side.

## Proposed Solution

### General

The interns from MCS-SPBU

- will work from begin of October until the end of December
- will spend eight working hours per calendar week working on the tasks

Truvity B.V.

- will provide the technical infrastructure to solve the technical tasks
- will provide necessary technical documentation and information to solve the technical tasks
- will provide once per calendar week 1-2 hours of consultation meeting with the interns (possible languages: Russian, English, in some cases English-only)
- will provide feedback about intermediate results and helps interns to figure out the next steps 

The task: build up the mathematical model and technical solution to manage legal documents

### Timeline

"Checkpoint" in milestone, which the intern should achieve.

Without achieving some milestone, an intern could not move forward.

For successful completion of the task, students need to complete all milestones.

#### 1. Introduction

Oleg Tsarev will introduce the company, business domain, and general idea of interns' activities in a short (1 hour) online interview
Oleg Tsarev will provide [golang tutorial](https://go.dev/doc/tutorial/getting-started)

Based on discussion with interns we agree about test task assignment:
- implement [Dinic's algoithm](https://en.wikipedia.org/wiki/Dinic%27s_algorithm)
- programming language: golang
- time box: approximately 4 hours
- result should be presented as github repository
- link to github you should send to email `mcs-spbu@project-plato.com`
- you can add additional information to your email (in addition to link) - like "why you choose this project" or "what you want to learn from this project"
- you need to send your email before September 28
- you will be informed about the result about September 30-October 3
- you could ask any additional question by email or Dmitry Shalymov

#### 2. Evaluation

Interns will provide technical task assignment solutions.
Oleg Tsarev will review technical task assignments and decide about "winners."

#### 3. Learning package

Oleg Tsarev will provide the learning materials.

- architecture meeting materials (except video) [legal-tech 101](../../architecture/meetings/2022-04-20/README.md)
- architecture decision record [ADR-GLO-01 | Legal Tech: common definitions](../../architecture/adr/adr-glo-01/README.md)
- [W3C: Verifiable Credentials Data Model](https://www.w3.org/TR/vc-data-model/)
- [JSON-LD](https://json-ld.org/learn.html) - Introductory Material - first two articles
- [golang tutorial](https://go.dev/doc/tutorial/getting-started)
- [Amazon Neptune: Getting Started](https://docs.aws.amazon.com/neptune/latest/userguide/graph-get-started.html)
- [Matrix Graph Grammar](https://repositorio.uam.es/bitstream/handle/10486/1815/5599_perez_velasco_pedro.pdf) by Pedro Pablo Perez Velasco

#### 4. Onboarding (checkpoint 1)

Oleg Tsarev will provide a manual to access the Amazon Neptune installation inside Truvity B.V. infrastructure.
Interns should follow the manual and solve simple graph tasks (load simple graph to Neptune and build up traversals based on this graph).

1. connect to Neptune
1. make simple read query
1. make simple write query

#### 5. Dealing with JSON-LD documents (checkpoint 2)

Oleg Tsarev will provide

- archive with JSON-LD documents (the result of e2e tests of hr-case)
- VChain diagram to illustrate schemas structure

Interns need to write the application to parse JSON-LD documents.

1. parse several files with documents with known schemas
2. print specific fields from parsed documents 

#### 6. Loading JSON-LD documents to Neptune (checkpoint 3)

Interns need to find the suitable representation of JSON-LD documents in the graph database (Neptune)

1. decide on representation of documents in the graph database
1. parse JSON-LD documents
1. write documents to graph database according to decided structure

#### 7. Querying documents in Neptune (checkpoint 4)

Interns need to find a suitable way to query data from Neptune.

1. load to database EmploymentAgreement + CandidateAgreement + CandidatePassport documents
2. query list of EmploymentAgreement documents (IDs of the documents + Position)
3. query list of employess (Position from EmploymentAgreement, First Name from CandidatePassport, Last Name from CandidatePassport)

#### 8. Updating graph representations by new documents (checkpoint 5)

Interns should find a suitable way to update graph representation by new documents (Employment Addendum, Passport).

1. load to database bunch of EmploymentAddendum
2. load to database bunch of new passport for employees
3. adjust query from previous checkpoint to updated structure

#### 9. Summarizing the result (checkpoint 6)

Interns should provide small essays with recommendations on how to deal with JSON-LD documents in Amazon Neptune (summary from previous steps)

#### 10. Applying results to work

Based on the interns' work, Oleg Tsarev will summarize the result of the initiative in the report.

Based on this report, Oleg Tsarev will probably adjust existing ADRs (architecture decisions records) or write a new one.

## Follow-up Tasks

