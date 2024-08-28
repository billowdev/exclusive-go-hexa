# Temporal Workflow

Temporal คือ workflow orchestration & long run process ที่มี UI ให้เราสามารถดู log การทำงานของ workflow ต่าง ๆ



ถ้าเปรียบเทียบ Temporal กับ Celery: Temporal สามารถทำทุกอย่างที่ Celery ทำได้ แต่สิ่งที่แตกต่างกันคือใน Celery เราใช้ Redis หรือ RabbitMQ เป็นตัว broker สำหรับจัดการข้อความระหว่าง worker และ task.


ในทางกลับกัน เมื่อใช้ Temporal เราจะมี Temporal server ที่ทำหน้าที่เป็น orchestration workflow server ซึ่งทำหน้าที่เป็นตัวกลางในการจัดการ Workflow ทั้งหมด ไม่ว่าจะเป็นการ Retries, การ Cancel workflow, และการจัดการสถานะของงาน เป็นต้น.


Qucik summary

  1. Worker ประกอบไปด้วย Workflow

  2. Workflow ประกอบไปด้วย Activity

  3. Workflow สามารถเพิ่มได้เรื่อย ๆ ถ้า ID ไม่ซ้ำกัน

  4. Workflow ถ้า ID ซ้ำกัน workflow ตัวนั้นต้องทำงาน Completed ก่อนถึงจะ รัน Workflow ที่ ID ซ้ำกันได้

  5. Workflow ภายใน Worker จะรันแบบ Parallel

  6. temporal workflow & activities ที่ถูก register เป็น worker และรันใน service ของเรา 

      - หาก มี worker ตัวอื่นที่ push task ไปที่ temporal server ที่มี workflow หรือ activities เหมือนกันเลย แสดงว่า worker ตัวอื่น มีสิทธิ์ pull task ตัวนั้น ไป execute

      - Workers: หลายตัวสามารถถูก register ให้ handle workflows หรือ activities แบบเดียวกันได้

      - Task Queues: Temporal server จะใช้ task queue สำหรับเก็บ tasks (งาน) ที่ถูกสร้างขึ้นเมื่อ workflow หรือ activity ถูกเรียกใช้งาน

      - Polling: Workers ที่ถูก register กับ Temporal server จะทำการ pull tasks จาก task queue ที่เกี่ยวข้องกับ workflow หรือ activity ที่พวกเขาสามารถ handle ได้

  7. เมื่อ workflow หรือ activity ถูก execute ตัว worker ที่รับผิดชอบงานนั้นจะทำการประมวลผลและส่งข้อมูลกลับไปที่ Temporal server ผ่าน gRPC ซึ่ง Temporal server จะจัดการ orchestration ทั้งหมด รวมถึงการจัดการ state, retries, และอื่นๆ
   
* Temporal has a 50K history event length limit and 50MB size limit for each execution.
  * ref: https://medium.com/@qlong/guide-to-continueasnew-in-cadence-temporal-workflow-using-iwf-as-an-example-part-1-c24ae5266f07

How it work ?

read more : https://mikhail.io/2020/10/practical-approach-to-temporal-architecture/

Get started 

How to implement temporal + golang


Example

read more: https://medium.com/lyonas/efficient-workflow-in-go-with-temporal-io-repository-pattern-d66e6ac78e39

Read more

* Temporal Youtube
https://www.youtube.com/@Temporalio/featured
* Why the data or history on database of temporal that encrypt or unreadable: https://community.temporal.io/t/check-workflow-execution-history-in-database/6489
เพราะว่าข้อมูลต่าง ๆ เช่น history ของ workflow จะถูกเก็บในรูปแบบของ serialized protobufs 
* Temporal database example
https://www.restack.io/docs/temporal-knowledge-temporal-database-examples
* (Series) Efficient Workflows in Go with Temporal: Signals & Selectors
https://medium.com/lyonas/series-efficient-workflows-in-go-with-temporal-signals-selectors-ddd4bbc285d4

