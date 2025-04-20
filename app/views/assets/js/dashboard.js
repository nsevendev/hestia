/********************  PARTIE ADMIN  ********************/

/* --------- Gestion du sélecteur “URL / Fichier” dans le formulaire news --------- */
function handleLinkTypeChange () {
    const container = document.getElementById('linkInputContainerNews');
    const select    = document.getElementById('linkTypeNews');
    const selected  = select ? select.value : '';
  
    /* Ré‑initialise le conteneur */
    container && (container.innerHTML = '');
  
    if (!container) return;
  
    if (selected === 'file') {
      container.innerHTML = `
        <label for="link">Fichier à joindre :</label>
        <input type="file" name="link" id="link" accept="*/*">
      `;
    }
  
    if (selected === 'url') {
      container.innerHTML = `
        <label for="link">URL du lien :</label>
        <input type="text" name="link" id="link" placeholder="https://...">
      `;
    }
  }
  
  /* Lance la détection (utile si le select possède déjà une valeur) */
  document.addEventListener('DOMContentLoaded', handleLinkTypeChange);
  
  
  
  /* =======================  SIDEBAR MOBILE  ======================= */
  
/* ============ SIDEBAR MOBILE ============ */
document.addEventListener('DOMContentLoaded', () => {
    const sidebar  = document.querySelector('.sidebar');
    const btnOpen  = document.getElementById('dashboard-nav-open');   // burger
    const btnClose = document.getElementById('dashboard-nav-close');  // croix
    const links    = sidebar ? sidebar.querySelectorAll('.sidebar-links a') : [];
  
    /* helpers */
    const open  = (e) => { e && e.preventDefault(); sidebar?.classList.add('show-sidebar'); };
    const close = (e) => {                 /* NE bloque *pas* les liens */
      if (e && (e.target === btnClose)) e.preventDefault();   // juste la croix
      sidebar?.classList.remove('show-sidebar');
    };
  
    btnOpen  && btnOpen .addEventListener('click', open);
    btnClose && btnClose.addEventListener('click', close);
    links.forEach(link => link.addEventListener('click', close)); // navigation conservée
  });