import { Component, OnInit } from '@angular/core';
import mockClassicData from '../data/card-data/mockClassicData.json'

interface DATAS {
  labels: string[];
  title: string;
  price: string;
  date: string;
  owner: string;
  srcImg: string;
}

@Component({
  selector: 'app-classic-product-cart',
  templateUrl: './classic-product-cart.component.html',
  styleUrls: ['./classic-product-cart.component.scss']
})
export class ClassicProductCartComponent implements OnInit {
  datas: DATAS[] = mockClassicData
  constructor() { }

  ngOnInit(): void {
    console.log(mockClassicData)
  }

}
