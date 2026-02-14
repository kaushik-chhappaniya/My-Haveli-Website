// js/directors.js

document.addEventListener("DOMContentLoaded", () => {
  loadDirectors();
});

async function loadDirectors() {
  try {
    const response = await fetch("data/directors.json");
    if (!response.ok) {
      throw new Error("Failed to fetch directors data");
    }

    const data = await response.json();
    renderDirectors(data.directors);
  } catch (error) {
    console.error(error);
  }
}

function renderDirectors(directors) {
  const tbody = document.getElementById("directors-table-body");
  tbody.innerHTML = "";

  directors.forEach((director) => {
    const row = document.createElement("tr");

    row.innerHTML = `
      <td>${director.srNo}</td>
      <td>${director.name}</td>
      <td>${director.role}</td>
      <td>${director.phone}</td>
    `;

    tbody.appendChild(row);
  });
}

function updateTimestamp(timestamp) {
  const el = document.getElementById("last-updated");
  el.textContent = `Last updated: ${timestamp}`;
}
