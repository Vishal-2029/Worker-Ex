## Worker Pattern with Nats Example :

- First Install The Nats for Windows :

  `https://github.com/nats-io/nats-server/releases/tag/v2.10.26`

- In code i connected to the nats with queue and worker. The queue is hold the 10 task and worker is received those task. Open multiple terminal and run worker.

 ## Running the Example 

 - First Start NATS Server :

 - Run Multiple Workers :
   
   ` go run main.go `

 - Run Queue :
   
   ` go run main.go `

 ## Outputs 

 - Queue Output :
   ```bash
       2025/03/19 15:32:20 Published: Task #1
       2025/03/19 15:32:20 Published: Task #2
       2025/03/19 15:32:20 Published: Task #3
       2025/03/19 15:32:20 Published: Task #4
       2025/03/19 15:32:20 Published: Task #5
       2025/03/19 15:32:20 Published: Task #6
       2025/03/19 15:32:20 Published: Task #7
       2025/03/19 15:32:20 Published: Task #8
       2025/03/19 15:32:20 Published: Task #9
       2025/03/19 15:32:20 Published: Task #10
       2025/03/19 15:32:20 All tasks published.
   ```

   - Workers Output :
     ```bash
               Worker 1                        |          Worker 1                          |           Worker 1
     -----------------------------------------------------------------------------------------------------------------
      2025/03/19 16:12:30 Received: Task #2    |  2025/03/19 16:12:30 Received: Task #1     |  2025/03/19 16:12:30 Received: Task #3
      2025/03/19 16:12:32 Completed: Task #2   |  2025/03/19 16:12:33 Completed: Task #1    |  2025/03/19 16:12:31 Completed: Task #3
      2025/03/19 16:12:32 Received: Task #5    |  2025/03/19 16:12:33 Received: Task #6     |  2025/03/19 16:12:31 Received: Task #4
      2025/03/19 16:12:34 Completed: Task #5   |  2025/03/19 16:12:36 Completed: Task #6    |  2025/03/19 16:12:34 Completed: Task #4
                                                                                            |  2025/03/19 16:12:34 Received: Task #7
                                                                                            |  2025/03/19 16:12:36 Completed: Task #7
                                                                                            |  2025/03/19 16:12:36 Received: Task #8
                                                                                            |  2025/03/19 16:12:39 Completed: Task #8
                                                                                            |  2025/03/19 16:12:39 Received: Task #9
                                                                                            |  2025/03/19 16:12:40 Completed: Task #9
                                                                                            |  2025/03/19 16:12:40 Received: Task #10
                                                                                            |  2025/03/19 16:12:42 Completed: Task #10
     
     ```
