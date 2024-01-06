
name="$QUERY_NAME"

if [[ -z "$name" ]]; then
    name="$(render-str "hello-form")"
fi

with-prop name "$name" |

render "homepage"
