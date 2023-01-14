#! /bin/bash

if [[ -z $1 ]];then 
    echo "param empty";exit;
fi
if [[ ! -d $1 ]];then 
    echo "not file";exit;
fi

PREFIX="master-"
ACCOUNT="sung1011"
ORIGIN="migrate"

SOURCE=$1
TARGET=$PREFIX$SOURCE

# cp -rf $SOURCE $TARGET

if [[ ! -z $2 ]];then 
    TARGET=$PREFIX$2
fi

cd $SOURCE;

cur_branch=$(git branch | grep \* |awk '{print $2}')
# git init;
git pull;
git add .;
git commit -m "base";
git remote add $ORIGIN git@github.com:$ACCOUNT/$TARGET.git;
git push -u $ORIGIN $cur_branch