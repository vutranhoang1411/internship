#include "iostream";

//wrong
class Employee{
public:
    Employee(){}
    void working(){

    }
    void checking(){

    }
};
class ParttimeWorker: public Employee{
    ParttimeWorker(){}
    void otherWork(){} //wrong, since Parttime worker can't check attendance, so ParttimeWorker can't inherit Employee
};

//right
class Employee2{
public:
    Employee2(){}
    void working(){

    }
};
class FullTimeEmployee:public Employee2{
public:
    FullTimeEmployee(){}
    void checking(){}
};
//now parttime Employee can inherit worker without any logic conflict
class ParttimeEmployee: public Employee2{
    ParttimeEmployee(){}
    void otherWork(){}
};