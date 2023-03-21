// Copyright (c) 2020 Mr. Coxall All rights reserved
//
// Created by: Charlotte Jhu
// Created on: March 2023
// This file contains the JS functions for index.html

/**
 * .* This function is called when the button is clicked
 */
function myButtonClicked() {
  // input
  const streetNumber = document.getElementById("street-number").value
  const streetName = document.getElementById("street-name").value

  // output
  document.getElementById("address").innerHTML = 
  "Your adress is " + streetNumber + ", " + streetName + "."
}
