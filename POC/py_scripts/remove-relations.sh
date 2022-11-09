
KIND=$(cat $1 | grep -e ^kind= | cut -d "=" -f2)
NAME=$(cat $1 | grep metadata.name= | cut -d "=" -f2)
if [[ ${#KIND} -eq 0 ]]; then
    echo -$KIND-
else
    python3 ./properties_to_yaml.py $1
    cat $1.yaml | grep -v 'id: ' | grep -v '^text_area_1: ' | grep -v '^name: ' | grep -v '^shape_library: ' | grep -v 'DRAFT' > $KIND-$NAME.yaml
    python3 replace.py --regex --in-place --from "('[0-9]': | '[0-9]':\n. *)" --to ' - ' $KIND-$NAME.yaml
    python3 replace.py --regex --in-place --from "apiversion" --to 'apiVersion' $KIND-$NAME.yaml
    python3 replace.py --regex --in-place --from "matchlabels" --to 'matchLabels' $KIND-$NAME.yaml
    python3 replace.py --regex --in-place --from "containerport" --to 'containerPort' $KIND-$NAME.yaml
    mv $KIND-$NAME.yaml $KIND-$NAME.yaml.bkp
    sed  -r "s/'([[:digit:]]+(\.[[:digit:]]+){0,1})'/\1/" $KIND-$NAME.yaml.bkp > $KIND-$NAME.yaml 
    rm -f $KIND-$NAME.yaml.bkp
fi