function scrollGridRight() {
	const grid = document.querySelector('.columns.is-mobile');
	const scrollAmount = grid.clientWidth / 2;
	
	grid.scrollBy({
		left: scrollAmount,
		behavior: 'smooth'
	});
}

function scrollGridLeft() {
	const grid = document.querySelector('.columns.is-mobile');
	const scrollAmount = grid.clientWidth / 2;
	
	grid.scrollBy({
		left: -scrollAmount,
		behavior: 'smooth'
	});
}