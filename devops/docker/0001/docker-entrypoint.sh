#!/usr/bin/env sh

CMD=$1

case "$CMD" in
  "empty" )
    echo "--- empty ---"

    exec /app/idogo
    ;;

   * )
    exec ${@}
    ;;
esac
