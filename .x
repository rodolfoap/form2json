touch .edit; MISSING=$(find . -type f -name \*.go|grep -v -f .edit); [ -z "${MISSING}" ] || echo "${MISSING}" >> .edit

execute(){
	rm -f app
	go build
	[ -f ./app ] || echo [ERROR] Compile ERROR.
	[ -f ./app ] && ./app
}
showjson(){
	[ -f $1 ] && { echo ===== $1 =====; cat $1|jq .; echo; echo; }
}

case "$1" in
e) 	vi -p $(grep -v '^#' .edit)
	execute;
	;;
c)	echo
	showjson defns/instance/test.json
	showjson defns/scenario/test.json
	;;
"")	execute
	;;
esac
