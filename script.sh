#!/bin/bash
Start() {
    if nc -zw1 google.com 443; then
        echo "we have connectivity"
    else
        echo "not have internet connection"
        return
    fi

    wget "https://assets.01-edu.org/stats-projects/math-skills" > /dev/null 2>&1
    echo "Testing file successfull downloaded"

    chmod +x math-skills
    go build cmd/main.go

    for i in {0..100} 
    do
        f1=$(./math-skills data.txt)
        f2=$(./main data.txt)
        if [[ "$f1" != "$f2" ]]; then
            echo "WRONG!!!"
            echo "$f1"
            echo "==================="
            echo "$f2"
            echo "\n"
            
        else
            echo "[$i] Success!"
        fi
    done
    rm math-skills
    echo "DONE!"
}

Start