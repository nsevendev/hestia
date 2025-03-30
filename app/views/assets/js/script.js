// Afficher / Masquer le logo dans la topbar selon la position du scroll (Hors rÃ©solution PC)

const logo = document.querySelector('.logo')

window.addEventListener('scroll', () => {
  if (window.scrollY > 100) {
    logo.classList.add('scrolled')
  } else {
    logo.classList.remove('scrolled')
  }
})

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
burgerIcon.addEventListener('click', function(event) {
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
    link.addEventListener('click', function() {
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

// Caroussel des avis client 

document.addEventListener('DOMContentLoaded', function() {
    const prev = document.querySelector('.carousel-control.prev');
    const next = document.querySelector('.carousel-control.next');
    const carousel = document.querySelector('.carousel');
    let scrollPosition = 0;
    const cardWidth = document.querySelector('.card-review').offsetWidth;
    const visibleCards = 3;
    
    prev.addEventListener('click', () => {
        scrollPosition = Math.max(scrollPosition - cardWidth * visibleCards, 0);
        carousel.style.transform = `translateX(-${scrollPosition}px)`;
    });
    
    next.addEventListener('click', () => {
        const maxScroll = (carousel.children.length - visibleCards) * cardWidth;
        scrollPosition = Math.min(scrollPosition + cardWidth * visibleCards, maxScroll);
        carousel.style.transform = `translateX(-${scrollPosition}px)`;
    });
});

// Bandeau temporaire 

document.addEventListener("DOMContentLoaded", function() {
  const startDate = new Date("2024-10-22T00:00:00+02:00"); 
  const endDate = new Date("2024-11-02T19:00:00+02:00");
  const now = new Date();
  const infoBanner = document.querySelector(".information");
  if (now >= startDate && now < endDate) {
      infoBanner.style.display = "block";
  } else {
      infoBanner.style.display = "none";
  }
});

/********************************* PARTIE ADMIN  ********************************/

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