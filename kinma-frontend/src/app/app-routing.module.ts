import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { MainPageComponent } from './main-page/main-page.component';
import { PageNotFoundComponent } from './page-not-found/page-not-found.component';

const routes: Routes = [
    {
      path: 'main', component: MainPageComponent,
      children: [{path: 'login', component:LoginComponent}]
    },
    {path: 'login', component:LoginComponent},
    {path: '', redirectTo: 'main', pathMatch: 'full'},
    {path: '**', component:PageNotFoundComponent},
  ];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
export const routingComponents = []