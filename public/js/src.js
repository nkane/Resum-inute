$(document).ready(function() {
	var videoFileUploadInput = $('#videoFileUploadInput');
	videoFileUploadInput.on("change", function(event) {
		if (event.target.files) {
			var videoFileUploadLabel = $('#videoFileUploadLabel');
			videoFile = event.target.files[0];
			videoFileUploadLabel[0].textContent = videoFile.name;
		}
	});

	var videoFileUploadButton = $('#videoFileUploadButton');
	videoFileUploadButton.on("click", function (event) {
		debugger;
		var formData = new FormData();
		var videoFileUploadInput = $('#videoFileUploadInput');
		formData.append('video', videoFileUploadInput[0].files[0]);
		$.ajax({
			url: 'video',
			type: 'POST',
			contentType: false,
			processData: false,
			data: formData,
			success: function(res) {
				console.log(res);
			},
			error: function(xhr, textStatus, error) {
				console.log(error);
			}
		});
	});
});
