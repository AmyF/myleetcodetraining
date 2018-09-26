# swiftCount=$(git ls-files | grep '.swift' -c)
# echo "$swiftCount"

files=$(git ls-files)
i=0
for line in $files
do
    file=$(basename $line)
    extArr[$i]=${file#*.}
    i=`expr $i + 1`
done

extArr=($(awk 'BEGIN{RS=" "} !a[$1]++' <<< ${extArr[@]}))

for ext in ${extArr[*]}
do
    count=$(git ls-files | grep ".$ext" -c)
    echo "$ext count is $count" 
done
