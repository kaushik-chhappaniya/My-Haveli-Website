fetch("data/notifications.json")
  .then((response) => {
    if (!response.ok) {
      throw new Error("Failed to load notifications");
    }
    return response.json();
  })
  .then((data) => {
    const container = document.getElementById("notificationsContainer");
    container.innerHTML = "";

    if (!data.notifications || data.notifications.length === 0) {
      container.innerHTML = `
        <div class="list-group-item text-muted text-center">
          No notifications available.
        </div>
      `;
      return;
    }

    data.notifications.forEach((n) => {
      const item = document.createElement("div");
      item.className = "list-group-item";

      item.innerHTML = `
        <div class="d-flex justify-content-between align-items-start">
           <div class="card-body table-hover">
      <h5 class="card-footer card-title  fw-bold">${n.title}</h5>
      <small class="text-muted">Updated -- ${n.timestamp}</small>
            <p class="mb-1 text-muted">${n.content}</p>
          </div>
        </div>
      `;

      container.appendChild(item);
    });
  })
  .catch((error) => {
    console.error(error);
    document.getElementById("notificationsContainer").innerHTML = `
      <div class="list-group-item text-danger text-center">
        Unable to load notifications.
      </div>
    `;
  });
