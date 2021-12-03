import { Component, OnInit } from '@angular/core';
import mockData from '../data/card-data/mockData.json'

interface DATAS {
  labels: string[];
  title: string;
  price: string;
  date: string;
  owner: string;
  srcImg: string;
}

@Component({
  selector: 'app-product-cart',
  templateUrl: './product-cart.component.html',
  styleUrls: ['./product-cart.component.scss']
})
export class ProductCartComponent implements OnInit {
  datas: DATAS[] = mockData;
  constructor() { }

  ngOnInit(): void {
  }

}
