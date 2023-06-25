const apiUri = "localhost:8000"

const groundsBtn = document.querySelector('#mapGrounds')
const weatherBtn = document.querySelector('#mapWeather')
const techBtn = document.querySelector('#mapTech')
const contentElem = document.querySelector('#content')

let currentRole

checkAuth()

async function checkAuth() {
    try {
        const response = await fetch('/auth/')
        const role = await response.json()
        currentRole = role
    } catch (error) {
        console.log(error)
        showSelectRole()
    }
}


function showSelectRole() {
    contentElem.innerHTML = roleTemplate.formatUnicorn({});

    document.forms['vibor'].addEventListener('submit', handleSubmitForm)

    async function handleSubmitForm(event) {
        event.preventDefault();

        const response = await fetch(event.target.action, {
            method: 'POST',
            body: new URLSearchParams(new FormData(event.target)) // event.target is the form
        })
        const data = await response.json()

        if (data.error) {
            document.querySelector('#error').textContent = data.error
            return
        }

        showWeatherMap()
    }
}

function showWeatherMap() {
    window.open("https://www.figma.com/proto/geF3iCHB49EXLRPwj6VVuN/Untitled?type=design&node-id=14-355&scaling=min-zoom&page-id=0%3A1&starting-point-node-id=14%3A355")
}




