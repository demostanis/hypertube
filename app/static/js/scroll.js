function updateArrowVisibility(categoryId) {
	if (categoryId instanceof Event)
		categoryId = categoryId.target.id;

	const grid = document.getElementById(categoryId);
	const category = grid.closest('.list');
	const leftArrow = category.querySelector('.arrow-left');
	const rightArrow = category.querySelector('.arrow-right');

	if (grid.scrollLeft <= grid.clientWidth / 7)
		leftArrow.style.opacity = 0;
	else
		leftArrow.style.opacity = 1;

	if (grid.scrollLeft >= (grid.scrollWidth - grid.clientWidth) - grid.clientWidth / 7)
		rightArrow.style.opacity = 0;
	else
		rightArrow.style.opacity = 1;
}

document.addEventListener('DOMContentLoaded', function() {
	const grids = document.querySelectorAll('.columns.is-mobile');
	grids.forEach(grid => {
		if (grid.id) {
			grid.addEventListener('scroll', function(e) {
				updateArrowVisibility(e.target.id);
			});
			updateArrowVisibility(grid.id);
		}
	});
});

function scrollGridRight(categoryId) {
	const grid = document.getElementById(categoryId);
	const scrollAmount = grid.clientWidth;

	grid.scrollBy({
		left: scrollAmount,
		behavior: 'smooth'
	});
}

function scrollGridLeft(categoryId) {
	const grid = document.getElementById(categoryId);
	const scrollAmount = grid.clientWidth;

	grid.scrollBy({
		left: -scrollAmount,
		behavior: 'smooth'
	});
}