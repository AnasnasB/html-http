async function requestpost(url, data) {
	let res = await fetch(url, {
		method: 'POST', // *GET, POST, PUT, DELETE, etc.
		mode: 'cors', // no-cors, *cors, same-origin
		cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
		credentials: 'same-origin', // include, *same-origin, omit
		headers: {
		  'Content-Type': 'application/json'
		  // 'Content-Type': 'application/x-www-form-urlencoded',
		},
		redirect: 'follow', // manual, *follow, error
		referrerPolicy: 'no-referrer', // no-referrer, *client
		bodypost: data // body data type must match "Content-Type" header
	  });
	let inf = await res;
	console.log(inf);
	
	let jso = await res.text();
	document.getElementById("resp").innerHTML = jso;
}

async function requestput(url, data) {
	let res = await fetch( url, {
		method: 'PUT',
		mode: 'cors',
		cache: 'no-cache',
		credentials: 'same-origin',
		headers: {
			'Content-Type': 'application/json'
		},
		redirect: 'follow',
		referrerPolicy: 'no-referrer',
		bodyput: data
	});
	let inf = await res;
	console.log(inf);
	
	let jso = await res.text();
	document.getElementById("resp").innerHTML = jso;	
		
}
	


function getpost() {
	
	requestpost(document.getElementById("url").value, document.getElementById("bodypost").value);
}

function getput() {
	
	requestput(document.getElementById("url").value, document.getElementById("bodyput").value);
}

function getread() {
	
	
}
