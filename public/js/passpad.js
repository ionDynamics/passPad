$(document).ready(function (){
	var trigger = undefined	
	$(".clipboard-trigger").hover(function(e) {
		//inHandler
		trigger = $(this)
	}, function(e) {
		//outHandler
		trigger = undefined
	})

	$(document).keydown(function(e) {
		if (!(e.ctrlKey || e.metaKey)) {
			return
		}

		if (getSelectionText() != "") {
			return
		}

		if (trigger) {
			clipboardContainer = $("#clipboard-container")
			clipboardContainer.empty().show()
			$("<textarea id='clipboard'></textarea>")
				.val(trigger.attr("x-data-tbc"))
				.appendTo(clipboardContainer)
				.focus()
				.select()							
		}					
	})


	$(document).keyup(function(e) {
		if ($(e.target).is("#clipboard")) {
			$("#clipboard-container").empty().hide()
		}
	})
})

function getSelectionText() {
	var text = "";
	if (window.getSelection) {
		text = window.getSelection().toString();
	} else if (document.selection && document.selection.type != "Control") {
		text = document.selection.createRange().text;
	}
	return text;
}