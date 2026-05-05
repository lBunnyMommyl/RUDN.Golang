document.addEventListener("DOMContentLoaded", function () {
    const servicesList = document.getElementById("servicesList");
    const requestForm = document.getElementById("requestForm");

    if (servicesList) {
        loadServices(servicesList);
    }

    if (requestForm) {
        requestForm.addEventListener("submit", sendMessage);
    }
});

async function loadServices(servicesList) {
    try {
        const response = await fetch("/api/services");

        if (!response.ok) {
            servicesList.innerHTML = '<p class="body3">Ошибка загрузки услуг</p>';
            return;
        }

        const services = await response.json();
        servicesList.innerHTML = "";

        services.forEach(function (service) {
            const card = document.createElement("div");
            card.className = "card card_content";

            card.innerHTML = `
                <p class="h3">${service.title}</p>
                <img src="images/${service.image}" class="card__img" alt="${service.title}">
                <p class="body3">Цена: ${service.price}<br>Срок: ${service.duration}</p>
                <a href="serv1.html" class="card_btn">
                    <p class="body3 text-white">Подробнее</p>
                </a>
            `;

            servicesList.appendChild(card);
        });
    } catch (error) {
        console.log(error);
        servicesList.innerHTML = '<p class="body3">Не удалось загрузить услуги</p>';
    }
}

async function sendMessage(event) {
    event.preventDefault();

    const data = {
        name: document.getElementById("name").value.trim(),
        email: document.getElementById("email").value.trim(),
        message: document.getElementById("message").value.trim()
    };

    if (data.name === "" || data.email === "" || data.message === "") {
        alert("Заполните все поля");
        return;
    }

    try {
        const response = await fetch("/api/messages", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(data)
        });

        if (!response.ok) {
            alert("Проверьте заполнение формы");
            return;
        }

        window.location.href = "/success.html";
    } catch (error) {
        console.log(error);
        alert("Ошибка отправки заявки");
    }
}