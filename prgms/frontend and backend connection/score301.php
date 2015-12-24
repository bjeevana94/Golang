<?php
$servername = "localhost";
$username = "root";
$password = "kala";
$dbname = "gre";

if(!(mysqli_connect($servername,$username,$password))) {
    die("Connection failed: " . $conn->connect_error);
} 

$sql = "SELECT id, cname FROM college WHERE score >= 301 and score<=315;
$result = $conn->query($sql);

if ($result->num_rows > 0) {
    // output data of each row
    while($row = $result->fetch_assoc()) {
        echo "id: " . $row["id"]. " - Name: " . $row["cname"] "<br>";
    }
} else {
    echo "0 results";
}
$conn->close();
?>