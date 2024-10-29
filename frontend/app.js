const availableTicketsElem = document.getElementById("availableTickets");

// Function to update available tickets display
function updateAvailableTickets() {
    availableTicketsElem.innerText = `Tickets Left: ${availableTickets}`;
}

document.getElementById("ticketForm").addEventListener("submit", async function(event) {
    event.preventDefault();

    const ticketData = {
        buyerName: document.getElementById("buyerName").value,
        numTickets: parseInt(document.getElementById("numTickets").value),
        email: document.getElementById("email").value,
        paymentMode: document.getElementById("paymentMode").value
    };

    const response = await fetch("/api/book-ticket", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(ticketData)
    });

    if (response.ok) {
        document.getElementById("confirmation").innerText = "Ticket booked successfully!";
        // Clear the input fields
        document.getElementById("buyerName").value = "";
        document.getElementById("numTickets").value = "";
        document.getElementById("email").value = "";
        document.getElementById("paymentMode").value = "";

        // Update available tickets
        availableTickets -= ticketData.numTickets;
        updateAvailableTickets();
    } else {
        document.getElementById("confirmation").innerText = "Failed to book ticket.";
    }
});

// Initially set the available tickets count
let availableTickets = 100; // Set this to your maximum available tickets
updateAvailableTickets();
