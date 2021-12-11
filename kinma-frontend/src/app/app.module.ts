import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { MDBBootstrapModule } from 'angular-bootstrap-md';
import { CarouselModule } from 'ngx-owl-carousel-o';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { HeaderComponent } from './header/header.component';
import { BannerComponent } from './banner/banner.component';
import { ProductCartComponent } from './product-cart/product-cart.component';
import { ProjectCardComponent } from './project-card/project-card.component';
import { CategoryComponent } from './category/category.component';
import { ClassicProductCartComponent } from './classic-product-cart/classic-product-cart.component';
import { BlockHeaderComponent } from './block-header/block-header.component';
@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    BannerComponent,
    ProductCartComponent,
    ProjectCardComponent,
    CategoryComponent,
    ClassicProductCartComponent,
    BlockHeaderComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    
    MDBBootstrapModule.forRoot(),
    BrowserAnimationsModule,
    CarouselModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
