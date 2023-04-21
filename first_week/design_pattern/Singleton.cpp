#include <string>
#include <iostream>
 
using namespace std;
 
class Database
{
private:
    //set constructor as private method so outsider can't create new Database instance
    static Database* mInstancePtr;
    Database(string name){}
 
public:
    //create static method to give access to the Database instance
    static Database* getInstance(string name)
    {
        //just create instance if the pointer is null
        if (mInstancePtr==NULL) {
            mInstancePtr = new Database(name);
        }
        return mInstancePtr;
    }
};