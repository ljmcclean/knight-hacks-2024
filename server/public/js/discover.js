
const projs = document.querySelector("#projs");

const skillInput = document.querySelector("#skill-inp");
const addSkillBtn = document.querySelector("#add-skill-btn");
addSkillBtn.addEventListener("click", function(){
    console.log(14);
});

const searchInput = document.querySelector("#discover-search");
const searchBtn = document.querySelector("#search-btn");

let skills = [];



const renderPosts = () => {
    fetch('/api/project/1')
    .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok ' + response.statusText);
            }
            return response.json(); // Parse the response as JSON
        })
        .then(data => {
            console.log(data); // Handle the data
        })
        .catch(error => {
            console.error('There was a problem with the fetch operation:', error);
        });
}

renderPosts();