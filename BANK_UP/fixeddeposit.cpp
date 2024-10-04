#include "fixeddeposit.h"
#include <iostream>
#include <unistd.h>
#include <cstdlib>

using namespace std;
namespace validation{
    bool agree(){
        cout<<"Agree to Open Account?\n Type [Y/y] for Yes\n Type [N/n] for No"<<endl;
        char yn;
        cin>>yn;
        if(yn=='Y' || yn=='y')
        {
            return true;
        }
        return false;
    }
}
using namespace validation;
FixedDeposit::FixedDeposit(UserAccount& obj,int amount,int duration)
{
    this->accountNumber=obj.accountNumber;
        this->FDAmount=amount;
        this->duration=duration;

        if(duration<=3)
        {
            interest=7;
            if(agree())
            {
                openFD(obj,amount,duration);
            }
        }
        else if(duration<=5)
        {
            interest=7.2;
            if(agree())
            {
                openFD(obj,amount,duration);
            }
        }
        else if(duration<=10)
        {
            interest=7.5;
            if(agree())
            {
                openFD(obj,amount,duration);
            }
        }
        else
        {
            interest=7.9;
            if(agree())
            {
                openFD(obj,amount,duration);
            }
        }
}
void FixedDeposit::openFD(UserAccount& obj,int amount,int duration){
        fdInterestRate=interest;
        if(obj.age>=60)
        {
            this->fdInterestRate+=0.25;
        }
        if(obj.gender==0)
        {
            this->fdInterestRate+=0.1;
        }
        cout<<"FD Opened for "<<obj.accountNumber<<endl;
        cout<<"Final Amount after "<<duration<<" Years: "<<amount+(amount*fdInterestRate*duration)/100<<endl;
        sleep(2);
        system("cls");
}

string FixedDeposit::getAccountNumber() const{
        return this->accountNumber;
    }

int FixedDeposit::getFAmount() const{
        return this->FDAmount;
    }

float FixedDeposit::getInterestRate() const{
        return this->fdInterestRate;
    }

