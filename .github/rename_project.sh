#!/usr/bin/env bash
while getopts a:n:u:d: flag
do
    case "${flag}" in
        a) author=${OPTARG};;
        n) name=${OPTARG};;
        u) urlname=${OPTARG};;
        d) description=${OPTARG};;
    esac
done

echo "Author: $author";
echo "Project Name: $name";
echo "Project URL name: $urlname";
echo "Description: $description";

echo "Renaming project..."

original_author="aceodus"
original_name="go_echo_boilerplate"
original_urlname="go-echo-boilerplate"
original_description="Awesome go_echo_boilerplate created by aceodus"
# for filename in $(find . -name "*.*") 
for filename in $(git ls-files) 
do
    if [[ $filename == .github/workflows/* ]]; then
        echo "Ignored $filename"
        continue
    fi
    sed -i "s/$original_author/$author/g" $filename
    sed -i "s/$original_name/$name/g" $filename
    sed -i "s/$original_urlname/$urlname/g" $filename
    sed -i "s/$original_description/$description/g" $filename
    echo "Renamed $filename"
done

mv go_echo_boilerplate $name

# This command runs only once on GHA!
rm -rf .github/template.yml