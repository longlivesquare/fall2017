<?php

abstract class Person {
    private $fName;
    private $lName;
    private $DOB;
    private $address;

    public function getFName(){return $this->fName;}
    public function getLName(){return $this->lName;}
    public function getDOB(){return $this->DOB;}
    public function getAddress(){return $this->address;}
    public function setFName($n) {$this->fName = $n;}
    public function setLName($n) {$this->lName = $n;}
    public function setDOB($d) {$this->DOB = $d;}
    public function setAddress($a) {$this->address = $a;}

    public function getAge() {
        $born = strtotime($DOB);
        $age = date_diff(date("Y/m/d"),$born);
    }

    abstract public function dispInfo();
    abstract public function getJson();
}

class Student extends Person{
    private $id;
    private $gpa;
    private $units;

    public function __construct() {
        $this->units = 0;
        $this->gpa = 0.0;
        $this->id = -1;
        $this->fName = "";
        $this->lName = "";
        $this->$DOB = "";
        $this->address = "";
    }

    public function getId(){return $this->id;}
    public function getGpa(){return $this->gpa;}
    public function getUnits(){return $this->units;}
    public function setId($i){$this->id = $i;}
    public function setGPA($g){$this->gpa = $g;}
    public function setUnits($u){$this->units = $u;}

    public function dispInfo() {

    }

    public function getJson() {

    }

    public function getInfoStr() {
        return $this->fName . " " . $this->lName . "; " . $this->DOB . ", " . $this->address . "";
    }
}

class Faculty {

}















?>