
const projs = document.querySelector("#projs");

const skillInput = document.querySelector("#skill-inp");
const addSkillBtn = document.querySelector("#add-skill-btn");
const skillList = document.querySelector("#profile-skill-list");

const searchInput = document.querySelector("#discover-search");
const searchBtn = document.querySelector("#search-btn");

let skills = [];

const addListItem = () => {
    if(skillInput.value !== ""){
        skills.push(skillInput.value);
    }
    skillInput.value = "";
    console.log("This is in add list item");
}

const renderSkills = () => {
    skillList.innerHTML = "";

    for(let i = 0; i < skills.length; i++){
        skillList.innerHTML += `<li>${skills[i]}</li>`;
    }
    console.log("this is in renderSkills")
}

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

addSkillBtn.addEventListener("click", function(){
    addListItem();
    renderSkills();
});

skillInput.addEventListener('keydown', function(event) {
    if (event.key === 'Enter') {
        addListItem();
        renderSkills();
    }
});