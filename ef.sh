ef() {
    DIR=$(findpath --find-args='-type d' "$@" 2> /dev/null)
    if [ "$?" -eq "0" ]; then
        cd $DIR
    else
        cd "$@"
    fi
}
