<h1 align="center">One stop for your Golang needs</h1>
<p align="center">Look through the other branches for more information</p>

<br><br>

<h2 align="center">Core concepts:</h2>
<h4>Variables:</h4>
<ol>
	<li>Short hand declaration</li>
	<li>Package level declaration</li>
	<li>Scopes</li>
	<li>Empty variables</li>
</ol>

<hr>

<h4>Composite literals:</h4>
<ol>
	<li>Slices</li>
	<li>Maps</li>
	<li>Structs</li>
</ol>

<hr>

<h4>Functions:</h4>
<ol>
	<li>Returns</li>
	<li>Receivers</li>
	<li>Composition</li>
	<li>Interface</li>
	<li>Polymorphism</li>
</ol>

<hr>

<br><br>

<h2 align="center">Conventions:</h2>

<ol>
	<li>Start your project by creating a module. Run "go mod init <your_desired_project_name>"
	<li>Use "main.go"  as the entry point to your project.</li>
	<li>All methods used on imports start with a capital letter.</li>
	<li>Short hand variable declaration such as ":=" cannot be used outside of the function.</li>
	<li>Formatting with "%q" only adds double quotes to string.</li>
	<li>Formatting with "%T" outputs the variables type.</li>
	<li>Formatting with "%f" is used for outputting floating point numbers.</li>
	<li>Formatting with "%0.1f" is used for outputting floating point numbers with a certain number of integers after the decimal point.</li>
	<li>Slices are mostly used when compared to Arrays.</li>
	<li>Objects need to be declared outside of the function in to order to be access by the other file.</li>
	<li>Receiver functions are usually all with pointers or without pointers</li>
</ol>