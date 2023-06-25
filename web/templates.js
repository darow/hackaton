String.prototype.formatUnicorn = String.prototype.formatUnicorn ||
    function () {
        "use strict";
        let str = this.toString();
        if (arguments.length) {
            let t = typeof arguments[0];
            let key;
            let args = ("string" === t || "number" === t) ?
                Array.prototype.slice.call(arguments)
                : arguments[0];

            for (key in args) {
                str = str.replace(new RegExp("\\{" + key + "\\}", "gi"), args[key]);
            }
        }

        return str;
    };

let roleTemplate = `
<div class="container">
    <form id="vibor" action="/session">
        <div class="form-group">
            <label class="mb-1" for="selectJob">Выберите профессию:</label>
            <select class="form-control mb-2" id="role" name="role">
                <option>агроном</option>
                <option>виноградарь</option>
            </select>
        </div>
        <button type="submit" class="btn btn-primary">Отправить запрос</button>
    </form>
</div>

<div id="error"></div>
`