merge into bonuses d 
   using (select employee_id.* from employees) s 
   on (employee_id = a) 
   when matched then update set d.bonus = bonus 
     delete where (salary > 8000)
   when not matched then insert (d.employee_id, d.bonus) 
     values (s.employee_id, s.salary)
     where (s.salary <= 8000);

merge into (select * from bonuses) d
  using (select employee_id.* from employees) s 
   on (employee_id = a) 
   when matched then update set d.bonus = bonus 
     delete where (salary > 8000)
   when not matched then insert (d.employee_id, d.bonus) 
     values (s.employee_id, s.salary)
     where (s.salary <= 8000);