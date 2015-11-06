<?php
	$url = "http://localhost:8080/shortlink/";
	
	$data = '{"url": "http://www.google.com"}';
	
	$options = array(
		'http' => array(
			'header'=> "Content-Type: application/json\r\n" .
				"Accept: application/json\r\n",
			'method'  => 'POST',
			'content' => $data,
		),
	);
	$context = stream_context_create($options);
	$result = file_get_contents($url, false, $context);

	var_dump($result);
?>
