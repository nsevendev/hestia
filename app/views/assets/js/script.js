// Afficher les actualités 

let currentNewsCount = 4; // Nombre d'actualités affichées par défaut
const newsPerPage = 4; // Nombre d'actualités à afficher à chaque clic sur "Voir plus"

// Initialiser l'affichage des 4 premières actualités
document.addEventListener("DOMContentLoaded", function() {
    showNews(currentNewsCount);
});

// Fonction pour afficher les actualités
function showNews(count) {
    const allNews = document.querySelectorAll('.news');
    for (let i = 0; i < count; i++) {
        if (allNews[i]) {
            allNews[i].style.display = 'flex'; // Afficher l'actualité
        }
    }
}

// Bouton pour activer l'audio 

function toggleSound() {
  const video = document.querySelector('.hero-video');
  video.muted = !video.muted;

  const btn = document.querySelector('.sound-toggle');
  const icon = btn.querySelector('i');

  if (video.muted) {
    icon.classList.remove('fa-volume-up', 'fa-volume-high', 'fa-volume');
    icon.classList.add('fa-volume-xmark'); // Icône muet
  } else {
    icon.classList.remove('fa-volume-xmark');
    icon.classList.add('fa-volume-up'); // Icône son activé
  }
}

// ATTACHER l'événement clic ici :
document.querySelector('.sound-toggle').addEventListener('click', toggleSound);


// Écouteur d'événements pour le bouton "Voir plus d'actualités"
document.getElementById('load-more').addEventListener('click', function() {
    currentNewsCount += newsPerPage; // Incrémenter le nombre d'actualités affichées
    showNews(currentNewsCount); // Afficher plus d'actualités
});


// Afficher / Masquer le logo dans la topbar selon la position du scroll (Hors rÃ©solution PC)

const nav = document.querySelector('nav');

window.addEventListener('scroll', () => {
  const nav = document.querySelector('nav');
  const scrollY = window.scrollY;
  const maxScroll = 300; // Jusqu'à 300px, l'opacité va de 0.5 à 1

  let opacity = 0.5 + (scrollY / maxScroll) * 0.5;
  if (opacity > 1) opacity = 1;

  nav.style.backgroundColor = `rgba(0, 0, 0, ${opacity})`;
});

window.addEventListener('DOMContentLoaded', () => {
  const navLinks = document.querySelectorAll('.nav-link');

  function updateActiveLink() {
    let scrollPosition = window.scrollY;

    navLinks.forEach(link => {
      const target = document.querySelector(link.getAttribute('href'));
      if (target) {
        const targetTop = target.offsetTop - document.querySelector('nav').offsetHeight;
        const targetBottom = targetTop + target.offsetHeight;
        if (scrollPosition >= targetTop && scrollPosition < targetBottom) {
          navLinks.forEach(link => {
            link.classList.remove('active');
          });
          link.classList.add('active');
        } else {
          link.classList.remove('active');
        }
      }
    });
  }

  window.addEventListener('scroll', updateActiveLink);

  updateActiveLink();
});

// Sélection du bouton burger et de la barre de navigation
const burgerIcon = document.getElementById('burger-icon');
const navbar = document.getElementById('navbar');

// Fonction pour afficher ou masquer la barre de navigation
function toggleNavbar() {
  navbar.classList.toggle('show-navbar');
  burgerIcon.querySelector('i').classList.toggle('fa-bars');
  burgerIcon.querySelector('i').classList.toggle('fa-times');
}

// Ajouter un événement de clic au bouton burger
burgerIcon.addEventListener('click', function (event) {
  event.preventDefault(); // Empêche le comportement par défaut du lien
  toggleNavbar();
});

const navLinks = document.querySelectorAll('.nav-link');

// Fonction pour fermer la barre de navigation
function closeNavbar() {
  navbar.classList.remove('show-navbar');
  burgerIcon.querySelector('i').classList.remove('fa-times');
  burgerIcon.querySelector('i').classList.add('fa-bars');
}

// Ajouter un gestionnaire d'événement de clic à chaque lien de navigation
navLinks.forEach(link => {
  link.addEventListener('click', function () {
    closeNavbar();
  });
});

// Ajouter un gestionnaire d'événement de clic à chaque lien social
const socialLinks = document.querySelectorAll('.navbar .social-footer a');

socialLinks.forEach(link => {
  link.addEventListener('click', function () {
    closeNavbar();
  });
});

// GALLERY //

let currentIndex = 0; // Index de l'image actuellement affichée
const slides = document.querySelectorAll('.gallery-slides img'); // Sélectionne toutes les images
const totalSlides = slides.length; // Nombre total d'images

function showSlide(index) {
  // S'assurer que l'index est dans les limites
  if (index < 0) {
    currentIndex = totalSlides - 1; // Si l'index est inférieur à 0, aller à la dernière image
  } else if (index >= totalSlides) {
    currentIndex = 0; // Si l'index dépasse la limite, revenir à la première image
  } else {
    currentIndex = index;
  }

  // Masquer toutes les images
  slides.forEach((slide) => (slide.style.display = 'none'));

  // Afficher l'image actuelle
  slides[currentIndex].style.display = 'block';
}

function prevSlide() {
  showSlide(currentIndex - 1); // Afficher l'image précédente
}

function nextSlide() {
  showSlide(currentIndex + 1); // Afficher l'image suivante
}

// Afficher la première image au chargement initial
showSlide(currentIndex);

// FIN GALLERY //

// Mentions légales 

const legalLink = document.querySelector('footer a[href="#"]');
const legalModal = document.getElementById('legal-modal');

function openLegal() {
  legalModal.style.display = 'flex';
}

function closeLegal() {
  legalModal.style.display = 'none';
}

legalLink.addEventListener('click', function(event) {
  event.preventDefault();  
  openLegal();
});

document.querySelector('.close-legal').addEventListener('click', closeLegal);

//********************************* PARTIE ADMIN  ********************************/

function handleLinkTypeChange() {
  const container = document.getElementById("linkInputContainerNews");
  const selected = document.getElementById("linkTypeNews") === null ? "" : document.getElementById("linkTypeNews").value;

  if (container) {
    while (container.firstChild) {
      container.removeChild(container.firstChild)
    }
  }

  if (selected === "file") {
    const label = document.createElement("label")
    label.setAttribute("for", "link")
    label.textContent = "Fichier à joindre :"

    const input = document.createElement("input")
    input.type = "file"
    input.name = "link"
    input.id = "link"
    input.accept = "*/*"

    container.appendChild(label)
    container.appendChild(input)
  }

  if (selected === "url") {
    const label = document.createElement("label")
    label.setAttribute("for", "link")
    label.textContent = "URL du lien :"

    const input = document.createElement("input")
    input.type = "text"
    input.name = "link"
    input.id = "link"
    input.placeholder = "https://..."

    container.appendChild(label)
    container.appendChild(input)
  }
}

// au cas où il y aurait une valeur pré-sélectionnée à l'ouverture
document.addEventListener("DOMContentLoaded", handleLinkTypeChange);