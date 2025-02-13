function Factorial(n)
{	
    var fact = 1;
    for (i=2; i<=n; i++)
        fact *= i;
    return fact;
}
function Perm(m, n)
{	
    var perm = m;
    for (i=m-1; i>m-n; i--)
        perm *= i;
    return perm;
}
function Comb(m, n)
{	
    var result = Factorial(m) / (Factorial(n) * Factorial(m - n)); // Equal to -> Perm(n,r) / Factorial(r)
    return result;
}	

function Fibonacci(n){
    if(n == 0 || n == 1){
        return n;
    }
    return Fibonacci(n-1) + Fibonacci(n-2);
}

function Lucas(n){
    if(n == 0)
        return 2;
    else if(n == 1)
        return 1;
    else{
        return Lucas(n-1) + Lucas(n-2);
    }
}

function Eulerian(m, k){
    if(k > m || m == 0)
        return 0;
    else if(k == 0)
        return 1;
    else{
        return (m-k)*Eulerian(m-1, k-1) + (k+1)*Eulerian(m-1, k);
    }
}

function Onto(m, n){
    var sum = 0;
    for(var i = 0; i <= n; i++){
        sum += Math.pow(-1, i)*Comb(n, n-i)*Math.pow(n-i, m);
    }
    return sum;
}

function invalid(m, n)
{	
    if (parseInt(m) < 0 || isNaN(parseInt(n)))
    {	document.getElementById("m").focus();
        return true;
    }
    else if (parseInt(n) < 0 || isNaN(parseInt(n)))
    {	document.getElementById("n").focus();
        return true;
    }
    return false;
}

function computeAll()
{	
    document.getElementById("resPermutation").innerHTML = "";
    document.getElementById("resCombination").innerHTML = "";
    document.getElementById("catalan_result").innerHTML = "";
    document.getElementById("triangular_result").innerHTML = "";
    document.getElementById("harmonic_result").innerHTML = "";
    document.getElementById("fibonacci_result").innerHTML = "";
    document.getElementById("lucas_result").innerHTML = "";
    document.getElementById("eulerian_result").innerHTML = "";
    clickPerm();
    clickComb();
    click_catalan();
    click_triangular(); 
    click_harmonic();
    click_fibonacci();
    click_lucas();
    click_eulerian();
}

function compute_all_bonus(){
    document.getElementById("resArrangement").innerHTML = "";
    document.getElementById("onto_result").innerHTML = "";
    document.getElementById("stirling_with_empty_result").innerHTML = "";
    document.getElementById("stirling_result").innerHTML = "";
    document.getElementById("resSelectrofn").innerHTML = "";
    document.getElementById("combination_repetition_no_empty_result").innerHTML = "";
    document.getElementById("partition_with_empty_result").innerHTML = "";
    document.getElementById("partition_result").innerHTML = "";
    clickArra();
    click_onto();
    click_stirling_with_empty();
    click_stirling();
    clickSele();
    click_combination_repetition_no_empty();
    click_partition_with_empty();
    click_partiton();
}

function clickPerm()
{	
    document.getElementById("resPermutation").innerHTML = "";
    var m = document.getElementById("m").value;
    var n = document.getElementById("n").value;
    if (invalid(m, n)) return;
    document.getElementById("resPermutation").innerHTML = Perm(m, n);
}

function clickComb()
{	
    document.getElementById("resCombination").innerHTML = "";
    var m = document.getElementById("m").value;
    var n = document.getElementById("n").value;										
    if (invalid(m, n)) return;				
    document.getElementById("resCombination").innerHTML = Comb(m, n);
}

function clickArra()
{	
    document.getElementById("resArrangement").innerHTML = ""
    var m = document.getElementById("m2").value;
    var n = document.getElementById("n2").value;
    if (invalid(m, n)) {
        document.getElementById("resArrangement").innerHTML = "m or n must greater or equal to 0 !"
        return;
    }				
    document.getElementById("resArrangement").innerHTML = Math.pow(m, n);
}

function clickSele()
{	
    document.getElementById("resSelectrofn").innerHTML = "";
    var m = parseInt(document.getElementById("m2").value);
    var n = parseInt(document.getElementById("n2").value);
    if (invalid(m, n)) {
        document.getElementById("resSelectrofn").innerHTML = "m or n must greater or equal to 0 !"
        return;
    }
    document.getElementById("resSelectrofn").innerHTML = Comb(m + n - 1, n);
}

function click_catalan(){
    var catalan_result_element = document.getElementById("catalan_result");
    catalan_result_element.innerHTML = "";
    var n_element = document.getElementById("n");
    var n = parseInt(n_element.value);
    if(n < 0){
        catalan_result_element.innerHTML = "n must greater or equal to 0 !";
        n_element.focus();
        return;
    }
    var result = "";
    for(var i = 0; i < n; i++){
        result += 1/(1+i) * Comb(2*i, i);
        if(i != n-1){
            result += ", ";
        }
    }
    catalan_result_element.innerHTML = result;
}

function click_triangular(){
    var triangular_element = document.getElementById("triangular_result");
    triangular_element.innerHTML = "";
    var n_element = document.getElementById("n");
    var n = parseInt(n_element.value);
    if(n < 0){
        triangular_element.innerHTML = "n must greater or equal to 0 !";
        n_element.focus();
        return;
    }
    var result = "";
    for(var i = 0; i < n; i++){
        result += (i*(i+1))/2;
        if(i != n-1){
            result += ', ';
        }
    }
    triangular_element.innerHTML = result;
}

function click_harmonic(){
    var harmonic_element = document.getElementById("harmonic_result");
    harmonic_element.innerHTML = "";
    var n_element = document.getElementById("n");
    var n = parseInt(n_element.value);
    if(n < 1){
        harmonic_element.innerHTML = "n must greater than 0 !";
        n_element.focus();
        return;
    }
    var result = "";
    var sum = 0;
    for(var i = 1; i <= n; i++){
        sum += 1/i;
        result += Math.round(sum * 10000) / 10000; // 因為Math.round()只能四捨五入到整數位，所以要除以10000讓他四捨五入到小數點後第四位
        if(i != n){
            result += ', ';
        }
    }
    harmonic_element.innerHTML = result;
}

function click_fibonacci(){
    var fibonacci_element = document.getElementById("fibonacci_result");
    fibonacci_element.innerHTML = "";
    var n_element = document.getElementById("n");
    var n = parseInt(n_element.value);
    if(n < 0){
        fibonacci_element.innerHTML = "n must greater or equal to 0 !";
        n_element.focus();
        return;
    }
    var result = "";
    for(var i = 0; i < n; i++){
        result += Fibonacci(i);
        if(i != n-1){
            result += ', ';
        }
    }
    fibonacci_element.innerHTML = result;
}

function click_lucas(){
    var lucas_element = document.getElementById("lucas_result");
    lucas_element.innerHTML = "";
    var n_element = document.getElementById("n");
    var n = parseInt(n_element.value);
    if(n < 0){
        lucas_element.innerHTML = "n must greater or equal to 0 !";
        n_element.focus();
        return;
    }
    var result = "";
    for(var i = 0; i < n; i++){
        result += Lucas(i);
        if(i != n-1){
            result += ', ';
        }
    }
    lucas_element.innerHTML = result;
}

function click_eulerian(){
    var eulerian_element = document.getElementById("eulerian_result");
    eulerian_element.innerHTML = "";
    var m_element = document.getElementById("m");
    var n_element = document.getElementById("n");
    var m = parseInt(m_element.value);
    var n = parseInt(n_element.value);
    if(invalid(m, n)){
        eulerian_element.innerHTML = "m or n must greater or equal to 0 !";
        return;
    }
    eulerian_element.innerHTML = Eulerian(m, n); // 遞迴的方式
}

function click_stirling(){
    var stirling_element = document.getElementById("stirling_result");
    stirling_element.innerHTML = "";
    var m_element = document.getElementById("m2");
    var n_element = document.getElementById("n2");
    var m = parseInt(m_element.value);
    var n = parseInt(n_element.value);
    if(m < 1 || n < 1){
        stirling_element.innerHTML = "m or n must greater or equal to 1 !"
        return;
    }
    stirling_element.innerHTML = Onto(m, n) / Factorial(n);
}

function click_onto(){
    var onto_element = document.getElementById("onto_result");
    onto_element.innerHTML = "";
    var m_element = document.getElementById("m2");
    var n_element = document.getElementById("n2");
    var m = parseInt(m_element.value);
    var n = parseInt(n_element.value);
    if(m < 1 || n < 1){
        onto_element.innerHTML = "m or n must greater or equal to 1 !";
        return;
    }
    onto_element.innerHTML = Onto(m, n);
}

function click_stirling_with_empty(){
    var stirling_with_empty_element = document.getElementById("stirling_with_empty_result");
    stirling_with_empty_element.innerHTML = "";
    var m_element = document.getElementById("m2");
    var n_element = document.getElementById("n2");
    var m = parseInt(m_element.value);
    var n = parseInt(n_element.value);
    if(m < 1 || n < 1){
        stirling_with_empty_element.innerHTML = "m or n must greater or equal to 1 !";
        return;
    }
    var sum = 0;
    for(let i = 1; i<=n; i++){
        sum += Onto(m, i) / Factorial(i);
    }
    stirling_with_empty_element.innerHTML = sum;
}

function click_combination_repetition_no_empty(){
    var combination_repetition_no_empty = document.getElementById("combination_repetition_no_empty_result");
    combination_repetition_no_empty.innerHTML = "";
    var m_element = document.getElementById("m2");
    var n_element = document.getElementById("n2");
    var m = parseInt(m_element.value);
    var n = parseInt(n_element.value);
    if(m < 1 || n < 1){
        combination_repetition_no_empty.innerHTML = "m or n must greater or equal to 1 !";
        return;
    }
    combination_repetition_no_empty.innerHTML = Comb(m-1, m-n);
}

function Partition(m, n){
    if(n == 1)
        return 1;
    else if(m == n)
        return 1;
    else if(m <= 0 || n <= 0)
        return 0;
    else if(m < n)
        return 0;
    else{
        return Partition(m-1, n-1) + Partition(m-n, n);
    }
}

function click_partiton(){
    var partiton_element = document.getElementById("partition_result");
    partiton_element.innerHTML = "";
    var m_element = document.getElementById("m2");
    var n_element = document.getElementById("n2");
    var m = parseInt(m_element.value);
    var n = parseInt(n_element.value);
    if(m < 1 || n < 1){
        partiton_element.innerHTML = "m or n must greater or equal to 1 !";
        return;
    }
    partiton_element.innerHTML = Partition(m, n);
}

function click_partition_with_empty(){
    var partition_with_empty_element = document.getElementById("partition_with_empty_result");
    partition_with_empty_element.innerHTML = "";
    var m_element = document.getElementById("m2");
    var n_element = document.getElementById("n2");
    var m = parseInt(m_element.value);
    var n = parseInt(n_element.value);
    if(m < 1 || n < 1){
        partition_with_empty_element.innerHTML = "m or n must greater or equal to 1 !";
        return;
    }

    var sum = 0; 
    for(var i = 1; i <= n; i++){
        sum += Partition(m, i);
    }
    partition_with_empty_element.innerHTML = sum;
}