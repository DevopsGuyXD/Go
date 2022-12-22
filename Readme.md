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

<hr>

<br><br>

<h2 align="center">Important things to remember:</h2>
<ol>
	<li>"var" keyword is used to create a variable without any value</li>
	<li>"Sprintf" saves the value but doesn't print it to the screen</li>
	<li>There are two types of values.</li>
		<ol>
			<li>Non-Pointer values: Original values are not updated by default, a copy is made with an updated value. Use * in order to update the original value</li>
				<ul>
					<li>Strings</li>
					<li>Ints</li>
					<li>Floats</li>
					<li>Booleans</li>
					<li>Arrays</li>
					<li>Structs</li>
				</ul>
			<li>Pointer Wrapper values: The original values are updated</li>
				<ul>
					<li>Slices</li>
					<li>Maps</li>
					<li>Functions</li>
				</ul>
		</ol>
	<li>Struct's are used as blueprints like constructors in java</li>
	<li>If the struct values are uppercase then they are accessible from outside the package</li>
	<li>If the struct values are lowercase then they are accessible only within the package</li>
	<li>"Receivers" are functions that are associated with Structs</li>
	<li>"*" is used for pointing to an object in memory</li>
	<li>"&" is to used for finding the location of the object in memory</li>
	<li>You can build the projects for different operating systems</li>
		<ul>
			<li>GOOS = "windows" go build</li>
			<li>GOOS = "linux" go build</li>
		</ul>
</ol>