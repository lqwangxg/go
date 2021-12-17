#git config --global user.email "lqwangxg@gmail.com"
#git config --global user.name "lqwangxg"
commitmsg=$1
branchname=$2

if [ -z $commitmsg ]; then
  commitmsg="auto commit"
fi

git add . 
git commit -a -m "$commitmsg"

if [ -z $commitmsg ]; then
    echo "dest branch is empty."
else
  curbranch=`git branch | grep -oe "$branchname" `
  if [ $? != 0 ]; then
    curbranch=`git branch -a | grep -oe "$branchname" `
    if [ $? != 0 ]; then
      git branch $branchname
    fi 
  fi
  git push -u origin $branchname
fi
