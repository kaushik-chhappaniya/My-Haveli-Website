document.addEventListener("DOMContentLoaded", () => {
  fetch("data/donations.json")
    .then((response) => {
      if (!response.ok) {
        throw new Error("Failed to load donors data");
      }
      return response.json();
    })
    .then((data) => {
      const donors = data.notifications || [];
      const tableBody = document.getElementById("donorsTableBody");

      if (!tableBody) return;

      tableBody.innerHTML = "";

      donors.forEach((donor, index) => {
        const name = donor.title || "-";
        const donation =
          donor.donation && donor.donation.trim() !== ""
            ? donor.donation
            : "*";
        const occasion =
          donor.occasion && donor.occasion.trim() !== ""
            ? donor.occasion
            : "-";
        const timestamp = donor.timestamp || "-";

        const row = document.createElement("tr");

        row.innerHTML = `
          <td>${index + 1}</td>
          <td>${name}</td>
          <td>${donation}</td>
          <td>${occasion}</td>
          <td>${timestamp}</td>
        `;

        tableBody.appendChild(row);
      });
    })
    .catch((error) => {
      console.error("Error loading donors:", error);
    });
});
