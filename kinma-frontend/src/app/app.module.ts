import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppComponent } from './app.component';
import { AppRoutingModule,routingComponents } from './app-routing.module';

import { MatDialogModule } from '@angular/material/dialog'
import { FormsModule }   from '@angular/forms';
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
import { FooterComponent } from './footer/footer.component';
import { MainPageComponent } from './main-page/main-page.component';
import { LoginComponent } from './login/login.component';
import { PageNotFoundComponent } from './page-not-found/page-not-found.component';
import { RegisterComponent } from './register/register.component';
import { SignInComponent } from './sign-in/sign-in.component'
import { AuthService } from './auth.service'
@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    BannerComponent,
    ProductCartComponent,
    ProjectCardComponent,
    CategoryComponent,
    ClassicProductCartComponent,
    BlockHeaderComponent,
    FooterComponent,
    routingComponents,
    MainPageComponent,
    LoginComponent,
    PageNotFoundComponent,
    RegisterComponent,
    SignInComponent,

  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    
    MDBBootstrapModule.forRoot(),
    BrowserAnimationsModule,
    CarouselModule,
    MatDialogModule,
    FormsModule
  ],
  providers: [AuthService],
  bootstrap: [AppComponent],
})
export class AppModule { }
