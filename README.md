**_AE test task_**

####Description:https://agileengine.bitbucket.io/fsNDJmGOAwqCpzZx/
####Recomnedations: https://agileengine.bitbucket.io/fsNDJmGOAwqCpzZx/api/

**Core tasks:**
We are looking to build a money accounting system. The application should be a web service. It should not do any real “transactional” work, just emulate the financial transactions logic (debit and credit).

We emulate debit and credit operations for a single user, so we always have just one financial account.

No security is required. So don't worry about authentication.

No real persistence is expected. Please don't invest time into DB integration.

Please avoid wasting time for complex project configuration. Use configuration from an existing project, if you have one, or use project skeleton generation tools for your technologies. Default configuration would be completely enough.

### AC's
    Service must store the account value of the single user.
    Service must be able to accept credit and debit financial transactions and update the account value correspondingly.
    Any transaction, which leads to negative amount within the system, should be refused. Please provide http response code, which you think suits best for this case.
    Application must store transactions history. Use in-memory storage. Pay attention that several transactions can be sent at the same time. The storage should be able to handle several transactions at the same time with concurrent access, where read transactions should not lock the storage and write transactions should lock both read and write operations.
    It is necessary to design REST API by your vision in the scope of this task. There are some API recommendations. Please use these recommendations as the minimal scope, to avoid wasting time for not-needed operations.
    In general, the service will be used programmatically via its RESTful API. For testing purposes Postman or any similar app can be used.
    It should be possible to launch project/projects by a single-line-command. Please provide README.md
    Target completion time is 3 hours. We would rather see what you were able to do in 3 hours than a full-blown application you’ve spent days implementing. Note that in addition to quality, time used is also factored into scoring the task.


### Delivery
    Source code.
    Binary versions of your applications that are ready to run. No build should be required.
    Readme.

### Run
    