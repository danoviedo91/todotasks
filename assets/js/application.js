require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");

window.addEventListener('DOMContentLoaded', function() {

    // MODAL WINDOW CONFIRM DELETION FEATURE
    trashBtns = document.querySelectorAll('.js-wwc-trash-btn')

    if (trashBtns.length != 0) {

        for (let i = 0; i < trashBtns.length; i++) {
            trashBtns[i].addEventListener("click", function(event) {

                let deleteUrl = trashBtns[i].dataset.delete; // Gets the data-delete string
                let deleteBtn = document.querySelector('.js-wwc-confirm-delete-btn');
                deleteBtn.setAttribute('href', deleteUrl); // Sets the data-delete string to the href
                
            });
        }
    }

});
