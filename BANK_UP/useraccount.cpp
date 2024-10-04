#include "useraccount.h"
#include <iostream>
#include <unistd.h>
#include <cstdlib>
using namespace std;

namespace validation{
    string accountSelection(){
        cout<<"Choose Account type:"<<endl;
        cout<<"1.SAVINGS"<<endl;
        cout<<"2.CURRENT"<<endl;
        int opt;
        cin>>opt;
        if(opt==2) return "CURRENT";
        return "SAVINGS";
    }
    string branchSelection(){
        cout<<"Choose Branch Loaction:"<<endl;
        cout<<"1.BANGALORE"<<endl;
        cout<<"2.BOKARO"<<endl;
        int opt;
        cin>>opt;
        if(opt==1) return "BANGALORE";
        return "BOKARO";
    }
    bool genderSelection(){
        cout<<"Choose Gender:"<<endl;
        cout<<"1.MALE"<<endl;
        cout<<"2.FEMALE"<<endl;
        int opt;
        cin>>opt;
        if(opt==1) return 1;
        return 0;
    }

    int initialDeposit(){
        int a;
        cout<<"Initial Deposit Amount:";
        cin>>a;
        return a;
    }
    string AccGen(const string& num){
        return "BSHDFC"+num;
    }
}
using namespace validation;

UserAccount::UserAccount(const string& mobileNumber,const string& name,int age)
{
        this->accountNumber=AccGen(mobileNumber);
        this->mobileNumber=mobileNumber;
        this->accountHolderName=name;
        this->accountType=accountSelection();
        this->branch=branchSelection();
        this->totalFDs=0;
        this->balance=initialDeposit();
        this->age=age;
        this->gender=genderSelection();
        this->interestRate=accountType=="SAVINGS" ? 4:0;
        cout<<"New Account Opened Congratulations "<<name<<endl;
        sleep(2);
        system("cls");
}
string UserAccount::getAccountName(){
    return this->accountHolderName;
}
string UserAccount::getAccountNumber(){
    return this->accountNumber;
}

string UserAccount::getMobileNumber(){
    return this->mobileNumber;
}


